package handlers

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/archive"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/rs/zerolog"
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
	path := filepath.Join("b0-temp", userId, projectSlug)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func createProjectFolder(userId, projectSlug string) error {
	path := filepath.Join("b0-temp", userId, projectSlug)

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

// removeProjectFolder removes the project folder
func removeProjectFolder(userId, projectSlug string) error {
	path := filepath.Join("b0-temp", userId, projectSlug)

	err := os.RemoveAll(path)

	if err != nil {
		return err
	}

	return nil
}

func setupProjectContents(userId, projectSlug string, code *agent.CodeGeneration) error {
	baseDir := filepath.Join("b0-temp", userId, projectSlug)

	// create all the files
	for _, file := range code.FileContents {
		// Create full path by joining base directory with file path
		fullPath := filepath.Join(baseDir, file.Filename)

		// Ensure the directory structure exists
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory structure: %v", err)
		}

		// Write the file
		if err := os.WriteFile(fullPath, []byte(file.Content), os.ModePerm); err != nil {
			return fmt.Errorf("failed to write file %s: %v", file.Filename, err)
		}
	}

	return nil
}

func cloneProjectToTar(userId, projectSlug string) (io.ReadCloser, error) {
	localPath := filepath.Join("b0-temp", userId, projectSlug)

	tar, err := archive.TarWithOptions(localPath, &archive.TarOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create tar archive: %v", err)
	}

	return tar, nil
}
