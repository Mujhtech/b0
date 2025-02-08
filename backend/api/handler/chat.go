package handler

import (
	"context"
	"net/http"

	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/mujhtech/b0/services"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Chat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	projectId, err := getProjectIdFromPath(r)

	if err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	endpointId := queryParamOrDefault(r, "endpoint", "")

	dst := new(dto.ChatRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	if dst.Model != "" {
		if _, err = agent.GetModel(dst.Model); err != nil {
			_ = response.BadRequest(w, r, err)
			return
		}
	}

	findProjectService := services.FindProjectService{
		ProjectID:   projectId,
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
	}

	project, err := findProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	var endpoint *models.Endpoint

	if endpointId != "" {

		findEndpointService := services.FindEndpointService{
			EndpointID:   endpointId,
			EndpointRepo: h.store.EndpointRepo,
			User:         session.User,
		}

		endpoint, err = findEndpointService.Run(ctx)

		if err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

	}

	// id := uuid.New().String()
	// name := "test.txt"
	// progress := 0

	// // Create background context for the goroutine
	// bgCtx := context.Background()

	if err = h.sse.Publish(ctx, project.ID, sse.EventTypeTaskStarted, sse.AgentData{
		Message: "b0 is working on your request...",
	}); err != nil {
		log.Printf("failed to publish task started event: %v", err)
	}

	// Use background context in goroutine
	bgCtx := context.Background()

	go func(ctx context.Context, project *models.Project, endpoint *models.Endpoint) {

		workflows, value, err := h.agent.GenerateWorkflow(ctx, project.Description.String, agent.WithModel(agent.ToModel(dst.Model)))

		if err != nil {
			if err = h.sse.Publish(ctx, project.ID, sse.EventTypeTaskFailed, sse.AgentData{
				Message: value,
				Error:   err.Error(),
			}); err != nil {
				log.Printf("failed to publish task failed event: %v", err)
			}
		}

		if err = h.sse.Publish(ctx, project.ID, sse.EventTypeTaskUpdate, sse.AgentData{
			Message:   value,
			Workflows: workflows,
		}); err != nil {
			log.Printf("failed to publish task updated event: %v", err)
		}

		// for progress < 100 {
		// 	progress += 5
		// 	if err = h.sse.Publish(ctx, appID, sse.EventTypeUploadProgress, UploadProgress{
		// 		FileID:   id,
		// 		Name:     name,
		// 		Status:   UploadProgressStatusUploading,
		// 		Progress: progress,
		// 	}); err != nil {
		// 		log.Printf("failed to publish upload progress event: %v", err)
		// 	}
		// 	time.Sleep(1 * time.Second)
		// 	if progress == 100 {
		// 		if err = h.sse.Publish(ctx, appID, sse.EventTypeUploadCompleted, UploadProgress{
		// 			FileID:   id,
		// 			Name:     name,
		// 			Status:   UploadProgressStatusCompleted,
		// 			Progress: progress,
		// 		}); err != nil {
		// 			log.Printf("failed to publish upload completed event: %v", err)
		// 		}
		// 	}
		// }
		if err = h.sse.Publish(ctx, project.ID, sse.EventTypeTaskCompleted, sse.AgentData{
			Message: "b0 has completed your request",
		}); err != nil {
			log.Printf("failed to publish task started event: %v", err)
		}
	}(bgCtx, project, endpoint)

	//_ = response.Ok(w, r, "file uploaded", nil)

	_ = response.Ok(w, r, "ok", nil)
}
