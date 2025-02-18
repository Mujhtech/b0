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
	"github.com/mujhtech/b0/internal/util"
	"github.com/rs/zerolog"
)

type UpdateWorkflowPayload struct {
	ProjectId  string `json:"project_id"`
	EndpointId string `json:"endpoint_id"`
	Prompt     string `json:"prompt"`
}

func HandleUpdateWorkflow(aesCfb encrypt.Encrypt, store *store.Store, agent *aa.Agent, event sse.Streamer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {

		rawPayload, err := aesCfb.Decrypt(string(t.Payload()))

		if err != nil {
			return err
		}

		var payload UpdateWorkflowPayload

		if err := util.UnmarshalJSON([]byte(rawPayload), &payload); err != nil {
			return err
		}

		project, err := store.ProjectRepo.FindProjectByID(ctx, payload.ProjectId)

		if err != nil {
			return err
		}

		endpoint, err := store.EndpointRepo.FindEndpointByID(ctx, payload.EndpointId)

		if err != nil {
			return err
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
			Prompt:    payload.Prompt,
			Workflows: endpoint.Workflows,
		}, aa.WithModel(aa.ToModel(project.Model.String)))

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

		err = store.EndpointRepo.UpdateEndpoint(ctx, endpoint.ID, &models.Endpoint{
			Workflows: workflows,
		})

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
		}); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to create AI usage")
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message:            "b0 has successfully updated your workflow, reloading...",
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
