package handlers

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"math/rand/v2"

	"github.com/docker/docker/pkg/archive"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/agent"
	secretmanager "github.com/mujhtech/b0/internal/pkg/secret_manager"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/mujhtech/b0/internal/util"
	"github.com/rs/zerolog"
)

var (
	TempFolder = filepath.Join(os.Getenv("HOME"), "dev", "b0-temp")
)

func sendEvent(ctx context.Context, projectID string, eventType sse.EventType, data AgentData, event sse.Streamer) {

	var errorMsg string

	switch eventType {
	case sse.EventTypeTaskStarted:
		errorMsg = sse.FailedToPublishTaskStartedEvent
	case sse.EventTypeTaskUpdate:
		errorMsg = sse.FailedToPublishTaskUpdatedEvent
	case sse.EventTypeTaskCompleted:
		errorMsg = sse.FailedToPublishTaskCompletedEvent
	case sse.EventTypeTaskFailed:
		errorMsg = sse.FailedToPublishTaskFailedEvent
	default:
		errorMsg = "unknown event type"
	}

	if err := event.Publish(ctx, projectID, eventType, data); err != nil {
		zerolog.Ctx(ctx).Error().Msgf("%s: %v", errorMsg, err)
	}
}

func checkIfProjectFolderExists(userId, projectSlug string) (bool, error) {
	// check if folder exists
	path := filepath.Join(TempFolder, userId, projectSlug)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func createProjectFolder(userId, projectSlug string) error {
	path := filepath.Join(TempFolder, userId, projectSlug)

	err := os.MkdirAll(path, os.ModePerm) // #nosec G301

	if err != nil {
		return err
	}

	return nil
}

// removeProjectFolder removes the project folder
func removeProjectFolder(userId, projectSlug string) error {
	path := filepath.Join(TempFolder, userId, projectSlug)

	err := os.RemoveAll(path)

	if err != nil {
		return err
	}

	return nil
}

func setupProjectContents(userId, projectSlug string, code *agent.CodeGeneration) error {
	baseDir := filepath.Join(TempFolder, userId, projectSlug)

	// create all the files
	for _, file := range code.FileContents {
		// Create full path by joining base directory with file path
		fullPath := filepath.Join(baseDir, file.Filename)

		// Ensure the directory structure exists
		dir := filepath.Dir(fullPath)
		err := os.MkdirAll(dir, os.ModePerm) // #nosec G301

		if err != nil {
			return fmt.Errorf("failed to create directory structure: %v", err)
		}

		// Write the file
		err = os.WriteFile(fullPath, []byte(file.Content), os.ModePerm) // #nosec G306

		if err != nil {
			return fmt.Errorf("failed to write file %s: %v", file.Filename, err)
		}
	}

	return nil
}

func cloneProjectToTar(userId, projectSlug string) (io.ReadCloser, error) {
	localPath := filepath.Join(TempFolder, userId, projectSlug)

	tar, err := archive.TarWithOptions(localPath, &archive.TarOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create tar archive: %v", err)
	}

	return tar, nil
}

func generatePort() string {
	r := rand.IntN(1999) // #nosec G404

	port := r + 5000

	return fmt.Sprintf("%d", port)
}

func checkUsageLimit(ctx context.Context, s *store.Store, project *models.Project) (*models.User, error) {

	user, err := s.UserRepo.FindUserByID(ctx, project.OwnerID)

	if err != nil {
		return nil, fmt.Errorf("failed to get project owner")
	}

	usageCount, err := s.AIUsageRepo.GetTotalUsage(ctx, store.TotalAIUsageFilter{
		OwnerID: user.ID,
		Range:   store.TotalAIUsageFilterRangeMonth,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get usage count")
	}

	if user.SubscriptionPlan == "free" && usageCount.TotalUsage >= 20 {
		return user, fmt.Errorf("you have reached the maximum number of requests for the current month")
	}

	if user.SubscriptionPlan == "starter" && usageCount.TotalUsage >= 50 {
		return user, fmt.Errorf("you have reached the maximum number of requests for the current month")
	}

	if user.SubscriptionPlan == "pro" && usageCount.TotalUsage >= 100 {
		return user, fmt.Errorf("you have reached the maximum number of requests for the current month, enable pay as you go to continue")
	}

	return user, nil
}

func GetEnvVars(ctx context.Context, secretManager secretmanager.SecretManager, projectId, endpointId string) ([]*dto.Secret, error) {

	secrets := []*dto.Secret{}

	secretId := projectId

	if endpointId != "" {
		secretId = fmt.Sprintf("%s_%s", secretId, endpointId)
	}

	secretName := fmt.Sprintf("projects/b0/%s/env-variables", secretId)

	secret, err := secretManager.GetSecret(ctx, secretName)

	if err != nil {
		return nil, err
	}

	if secret != nil {
		if err := util.UnmarshalJSON(secret, &secrets); err != nil {
			return nil, err
		}
	}

	return secrets, nil
}
