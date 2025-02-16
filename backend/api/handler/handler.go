package handler

import (
	"context"

	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/container"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/mujhtech/b0/job"
)

type Handler struct {
	cfg    *config.Config
	ctx    context.Context
	store  *store.Store
	cache  cache.Cache
	agent  *agent.Agent
	sse    sse.Streamer
	job    *job.Job
	docker *container.Container
}

func New(
	cfg *config.Config,
	ctx context.Context,
	store *store.Store,
	cache cache.Cache,
	agent *agent.Agent,
	sse sse.Streamer,
	job *job.Job,
	docker *container.Container,
) (*Handler, error) {

	return &Handler{
		cfg:    cfg,
		ctx:    ctx,
		store:  store,
		cache:  cache,
		agent:  agent,
		sse:    sse,
		job:    job,
		docker: docker,
	}, nil
}
