package handler

import (
	"net/http"

	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/internal/pkg/response"
)

func (h *Handler) GetUsage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	usage, err := h.store.AIUsageRepo.GetTotalUsageInCurrentMonth(ctx, session.User.ID)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "usage retrieved", usage)
}
