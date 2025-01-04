package auth

import (
	"context"
	"net/http"

	"github.com/mujhtech/b0/database/models"
)

type Auth interface {
	UserAuthenticate(r *http.Request) (*UserSession, error)
	CreateToken(
		ctx context.Context,
		value string,
		tokenType string,
	) (*models.Token, string, error)
}
