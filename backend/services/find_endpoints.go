package services

import (
	"context"

	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
)

type FindEndpointsService struct {
	EndpointRepo store.EndpointRepository
	User         *models.User
	Query        dto.GetEndpointQuery
}

func (f *FindEndpointsService) Run(ctx context.Context) ([]*models.Endpoint, error) {

	endpoints, err := f.EndpointRepo.FindEndpointByProjectID(ctx, f.Query.ProjectID)

	if err != nil {
		return nil, err
	}

	return endpoints, nil
}
