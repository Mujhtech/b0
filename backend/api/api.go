package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mujhtech/b0/api/handler"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database/store"
	"github.com/rs/zerolog/hlog"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Api struct {
	cfg     *config.Config
	handler *handler.Handler
	cache   cache.Cache
	store   *store.Store
}

func New(
	cfg *config.Config,
	ctx context.Context,
	store *store.Store,
	cache cache.Cache,
) (*Api, error) {

	h, err := handler.New(cfg, ctx, store, cache)
	if err != nil {
		return nil, fmt.Errorf("failed to create handler: %w", err)
	}

	return &Api{
		handler: h,
		cfg:     cfg,
		cache:   cache,
		store:   store,
	}, nil

}

func (a *Api) BuildRouter() *chi.Mux {
	router := chi.NewMux()

	router.Use(middleware.SetupRequestLog())
	router.Use(chiMiddleware.NoCache)
	router.Use(chiMiddleware.Recoverer)
	router.Use(hlog.URLHandler("http.url"))
	router.Use(hlog.MethodHandler("http.method"))
	router.Use(middleware.WriteRequestIDHeader())
	router.Use(middleware.HLogAccessLogHandler())
	router.Use(middleware.ApplyCORS(a.cfg))

	//
	router.Route("/api", func(r chi.Router) {
		// v1 route
		r.Route("/v1", func(r chi.Router) {})

		// platform route
		r.Route("/platform", func(r chi.Router) {

			r.Use(
				chiMiddleware.Maybe(middleware.RequiredUserAuth(a.cfg, a.store, a.cache), shouldAllowAuth),
			)

			// auth route
			r.Route("/auth", func(r chi.Router) {
				r.Get(fmt.Sprintf("/{%s}", handler.AuthProviderKey), a.handler.Authenticate)
				r.Get(fmt.Sprintf("/{%s}/callback", handler.AuthProviderKey), a.handler.AuthenticateCallback)
				r.Post(fmt.Sprintf("/{%s}/callback", handler.AuthProviderKey), a.handler.AuthenticateCallbackPost)
			})

			// features
			r.Get("/features", a.handler.GetFeatures)

			// user route
			r.Route("/user", func(r chi.Router) {
				r.Get("/", a.handler.GetUser)
			})

			// projects route
			r.Route("/projects", func(r chi.Router) {
				r.Get("/", a.handler.GetProjects)
				r.Post("/", a.handler.CreateProject)
				r.Put(fmt.Sprintf("/{%s}", handler.ProjectParamId), a.handler.UpdateProject)
			})

			// projects route
			r.Route("/endpoints", func(r chi.Router) {
				r.Get("/", a.handler.GetEndpoints)
				r.Post("/", a.handler.CreateEndpoint)
				r.Put(fmt.Sprintf("/{%s}", handler.EndpointParamId), a.handler.UpdateEndpoint)
			})

		})
	})

	return router
}

var guestRoutes = []string{
	"/auth/github",
	"/auth/github/callback",
	"/auth/google",
	"/auth/google/callback",
	"/features",
}

func shouldAllowAuth(r *http.Request) bool {

	for _, route := range guestRoutes {
		if strings.HasSuffix(r.URL.Path, route) {
			return false
		}
	}

	return true
}
