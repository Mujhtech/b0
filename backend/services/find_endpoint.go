package services

import (
	"context"

	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/errors"
)

type FindEndpointService struct {
	EndpointRepo store.EndpointRepository
	User         *models.User
	EndpointID   string
}

func (f *FindEndpointService) Run(ctx context.Context) (*models.Endpoint, error) {

	endpoint, err := f.EndpointRepo.FindEndpointByID(ctx, f.EndpointID)

	if err != nil {
		return nil, err
	}

	if endpoint.OwnerID != f.User.ID {
		return nil, errors.ErrNotAuthorized
	}

	return endpoint, nil
}
