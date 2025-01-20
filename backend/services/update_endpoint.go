package services

import (
	"context"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/errors"
)

type UpdateEndpointService struct {
	Body         *dto.CreateEndpointRequestDto
	ProjectRepo  store.ProjectRepository
	EndpointRepo store.EndpointRepository
	User         *models.User
	EndpointID   string
}

func (u *UpdateEndpointService) Run(ctx context.Context) (*models.Endpoint, error) {

	endpoint, err := u.EndpointRepo.FindEndpointByID(ctx, u.EndpointID)

	if err != nil {
		return nil, err
	}

	project, err := u.ProjectRepo.FindProjectByID(ctx, u.Body.ProjectID)

	if err != nil {
		return nil, err
	}

	if project.OwnerID != u.User.ID {
		return nil, errors.ErrNotAuthorized
	}

	if endpoint.OwnerID != u.User.ID {
		return nil, errors.ErrNotAuthorized
	}

	if endpoint.ProjectID != project.ID {
		return nil, errors.ErrNotAuthorized
	}

	if u.Body.Name != "" {
		endpoint.Name = u.Body.Name
	}

	if u.Body.Path != "" {
		endpoint.Path = u.Body.Path
	}

	if u.Body.Method != "" {
		endpoint.Metadata = u.Body.Method
	}

	if u.Body.Description != "" {
		endpoint.Description = null.NewString(u.Body.Description, true)
	}

	if u.Body.Status != "" && endpoint.Status != u.Body.Status {
		endpoint.Status = u.Body.Status
	}

	if endpoint.IsPublic != u.Body.IsPublic {
		endpoint.IsPublic = u.Body.IsPublic
	}

	err = u.EndpointRepo.UpdateEndpoint(ctx, endpoint)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}
