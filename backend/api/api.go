package api

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/mujhtech/b0/api/handler"
	"github.com/mujhtech/b0/config"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Api struct {
	cfg     *config.Config
	handler *handler.Handler
}

func New(
	cfg *config.Config,
	ctx context.Context,

) (*Api, error) {

	h, err := handler.New(cfg, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create handler: %w", err)
	}

	return &Api{
		handler: h,
		cfg:     cfg,
	}, nil

}

func (a *Api) BuildRouter() *chi.Mux {
	router := chi.NewMux()

	router.Use(chiMiddleware.NoCache)
	router.Use(chiMiddleware.Recoverer)

	//

	return router
}
