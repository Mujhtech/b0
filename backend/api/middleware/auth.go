package middleware

import (
	"context"
	"net/http"

	"github.com/mujhtech/b0/auth"
	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/response"
)

const (
	authSessionKey key = iota
)

func RequiredUserAuth(cfg *config.Config, store *store.Store, cache cache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			jwt := auth.NewJWTAuth(cfg, store.UserRepo, store.TokenRepo, cache)

			session, err := jwt.UserAuthenticate(r)

			if err != nil {
				_ = response.Unauthorized(w, r, err)
				return
			}

			if session == nil {
				_ = response.Unauthorized(w, r, err)
				return
			}

			next.ServeHTTP(w, r.WithContext(
				withAuthSession(ctx, session),
			))

		})

	}
}

func withAuthSession(ctx context.Context, session *auth.UserSession) context.Context {
	return context.WithValue(ctx, authSessionKey, session)
}

func GetAuthSession(ctx context.Context) (*auth.UserSession, bool) {
	v, ok := ctx.Value(authSessionKey).(*auth.UserSession)

	return v, ok && v != nil
}
