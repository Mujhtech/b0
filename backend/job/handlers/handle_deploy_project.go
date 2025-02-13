package handlers

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/mujhtech/b0/database/store"
	aa "github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/container"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
	"github.com/mujhtech/b0/internal/pkg/sse"
)

func HandleDeployProject(aesCfb encrypt.Encrypt, store *store.Store, agent *aa.Agent, event sse.Streamer, container *container.Container) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {

		projectId, err := aesCfb.Decrypt(string(t.Payload()))

		if err != nil {
			return err
		}

		project, err := store.ProjectRepo.FindProjectByID(ctx, projectId)

		if err != nil {
			return err
		}

		if err = event.Publish(ctx, project.ID, sse.EventTypeTaskStarted, AgentData{
			Message: "b0 is working on your request...",
		}); err != nil {
			log.Printf("failed to publish task started event: %v", err)
		}

		return nil
	}
}
