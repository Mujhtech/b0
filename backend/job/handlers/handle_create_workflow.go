package handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/hibiken/asynq"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	aa "github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/rs/zerolog"
)

type AgentData struct {
	Log                string         `json:"log,omitempty"`
	Message            string         `json:"message,omitempty"`
	Error              string         `json:"error,omitempty"`
	Workflows          []*aa.Workflow `json:"workflows,omitempty"`
	Deploying          bool           `json:"deploying,omitempty"`
	Code               interface{}    `json:"code,omitempty"`
	ShouldReloadWindow bool           `json:"should_reload_window,omitempty"`
}

func HandleCreateWorkflow(aesCfb encrypt.Encrypt, store *store.Store, agent *aa.Agent, event sse.Streamer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {

		projectId, err := aesCfb.Decrypt(string(t.Payload()))

		if err != nil {
			return err
		}

		project, err := store.ProjectRepo.FindProjectByID(ctx, projectId)

		if err != nil {
			return err
		}

		catalog, err := aa.GetModelCatalog(project.Model.String)

		if err != nil {
			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Error: err.Error(),
			}, event)
			return nil
		}

		if _, err := checkUsageLimit(ctx, store, project); err != nil {
			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Error: err.Error(),
			}, event)
			return nil
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskStarted, AgentData{
			Message: "b0 is working on your request...",
		}, event)

		workflows, agentToken, err := agent.GenerateWorkflow(ctx, aa.WorkflowGenerationOption{
			Prompt: project.Description.String,
		}, aa.WithModel(catalog.Model))

		if err != nil {
			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Message: agentToken.Output,
				Error:   err.Error(),
			}, event)

			return err
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskStarted, AgentData{
			Message: "b0 is currently generating your workflow...",
		}, event)

		zerolog.Ctx(ctx).Info().Msgf("workflows: %v", workflows)

		// pick first workflow
		requestWorkflow := workflows[0]

		endpoint := &models.Endpoint{
			ID:          uuid.New().String(),
			OwnerID:     project.OwnerID,
			ProjectID:   project.ID,
			Name:        requestWorkflow.Name,
			Description: null.NewString(requestWorkflow.Instruction, requestWorkflow.Instruction != ""),
			Path:        requestWorkflow.Url,
			Method:      models.EndpointMethod(requestWorkflow.Method),
			Workflows:   workflows,
			Metadata:    null.NewString("{}", true),
			IsPublic:    false,
			Status:      models.EndpointStatusDraft,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err = store.EndpointRepo.CreateEndpoint(ctx, endpoint)

		if err != nil {
			return err
		}

		//
		if err = store.AIUsageRepo.CreateAIUsage(ctx, &models.AIUsage{
			ID:          uuid.New().String(),
			ProjectID:   project.ID,
			EndpointID:  null.NewString(endpoint.ID, true),
			OwnerID:     project.OwnerID,
			Model:       project.Model.String,
			UsageType:   "workflow",
			InputToken:  agentToken.Input,
			OutputToken: agentToken.Output,
			IsPremium:   catalog.IsPremium,
		}); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to create AI usage")
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message:            "b0 has successfully generated your workflow, reloading...",
			Workflows:          workflows,
			ShouldReloadWindow: true,
		}, event)

		// delay 1 seconds
		time.Sleep(1 * time.Second)

		sendEvent(ctx, project.ID, sse.EventTypeTaskCompleted, AgentData{
			Message: "b0 has successfully generated your workflow",
		}, event)

		return nil
	}
}
