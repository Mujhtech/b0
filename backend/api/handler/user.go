package handler

import (
	"net/http"

	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/internal/pkg/response"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	_ = response.Ok(w, r, "user retrieved", session.User)
}
