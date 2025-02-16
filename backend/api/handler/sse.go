package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/mujhtech/b0/internal/pkg/sse"
	jobHandler "github.com/mujhtech/b0/job/handlers"
	"github.com/mujhtech/b0/services"
	"github.com/rs/zerolog/log"
)

func (h *Handler) ProjectEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	projectId, err := getProjectIdFromPath(r)

	if err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	// endpointId, := queryParamOrDefault(r, "endpoint", "")

	chEvents, chErr, sseCancel := h.sse.Subscribe(ctx, projectId)

	defer func() {
		err := sseCancel(ctx)
		if err != nil {
			log.Error().Err(err).Msg("failed to cancel sse subscription")
		}
	}()

	response.Stream(ctx, w, h.ctx.Done(), chEvents, chErr)
}

func (h *Handler) ProjectLog(w http.ResponseWriter, r *http.Request) {
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

	if project.ContainerID.String == "" {
		_ = response.InternalServerError(w, r, fmt.Errorf("project container id is empty"))
		return
	}

	con, err := h.docker.GetContainer(ctx, project.ContainerID.String)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	if con.State.Status != "running" {
		_ = response.InternalServerError(w, r, fmt.Errorf("container is not running"))
		return
	}

	chEvents, chErr, sseCancel := h.sse.Subscribe(ctx, project.ID+"-log")

	bgCtx := context.Background()

	go func(ctx context.Context, project *models.Project) {

		reader, err := h.docker.GetContainerLogsStream(ctx, project.ContainerID.String)

		if err != nil {
			log.Error().Err(err).Msg("failed to get container logs")
			return
		}

		defer reader.Close()

		if err = h.sse.Publish(ctx, project.ID+"-log", sse.EventTypeLogStarted, jobHandler.AgentData{
			Message: "b0 is streaming logs...",
		}); err != nil {
			log.Printf("failed to publish log started event: %v", err)
		}

		buf := make([]byte, 8*1024) // 8KB buffer

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := reader.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					}
					if err = h.sse.Publish(ctx, project.ID+"-log", sse.EventTypeLogFailed, jobHandler.AgentData{
						Message: "failed to read container logs",
						Error:   err.Error(),
					}); err != nil {
						log.Printf("failed to publish log event: %v", err)
					}
					return
				}

				if n > 0 {
					if err = h.sse.Publish(ctx, project.ID+"-log", sse.EventTypeLogUpdated, jobHandler.AgentData{
						Log: string(buf[:n]),
					}); err != nil {
						log.Printf("failed to publish log event: %v", err)
					}
				}
			}
		}

	}(bgCtx, project)

	defer func() {
		err := sseCancel(ctx)
		if err != nil {
			log.Error().Err(err).Msg("failed to cancel sse subscription")
		}
	}()

	response.Stream(ctx, w, h.ctx.Done(), chEvents, chErr)
}
