package handler

import (
	"net/http"

	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Event(w http.ResponseWriter, r *http.Request) {
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
