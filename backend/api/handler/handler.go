package handler

import (
	"context"

	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/billing/stripe"
	"github.com/mujhtech/b0/internal/pkg/container"
	secretmanager "github.com/mujhtech/b0/internal/pkg/secret_manager"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/mujhtech/b0/job"
)

type Handler struct {
	cfg           *config.Config
	ctx           context.Context
	store         *store.Store
	cache         cache.Cache
	agent         *agent.Agent
	sse           sse.Streamer
	job           *job.Job
	docker        *container.Container
	billing       stripe.Stripe
	secretManager secretmanager.SecretManager
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
	billing stripe.Stripe,
	secretManager secretmanager.SecretManager,
) (*Handler, error) {

	return &Handler{
		cfg:           cfg,
		ctx:           ctx,
		store:         store,
		cache:         cache,
		agent:         agent,
		sse:           sse,
		job:           job,
		docker:        docker,
		billing:       billing,
		secretManager: secretManager,
	}, nil
}
