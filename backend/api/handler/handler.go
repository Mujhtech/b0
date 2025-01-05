package handler

import (
	"context"

	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database/store"
)

type Handler struct {
	cfg   *config.Config
	ctx   context.Context
	store *store.Store
	cache cache.Cache
}

func New(
	cfg *config.Config,
	ctx context.Context,
	store *store.Store,
	cache cache.Cache,
) (*Handler, error) {

	return &Handler{
		cfg:   cfg,
		ctx:   ctx,
		store: store,
		cache: cache,
	}, nil
}
