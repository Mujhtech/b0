package handler

import (
	"context"

	"github.com/mujhtech/b0/config"
)

type Handler struct {
	cfg *config.Config
	ctx context.Context
}

func New(
	cfg *config.Config,
	ctx context.Context,
) (*Handler, error) {

	return &Handler{
		cfg: cfg,
		ctx: ctx,
	}, nil
}
