package store

import (
	"context"

	"github.com/mujhtech/b0/database/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByID(ctx context.Context, id string) (*models.User, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, token *models.Token) error
	FindTokenByID(ctx context.Context, id string) (*models.Token, error)
	DeleteToken(ctx context.Context, id string) error
}
