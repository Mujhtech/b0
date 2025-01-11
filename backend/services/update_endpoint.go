package services

import (
	"context"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/errors"
	"github.com/mujhtech/b0/internal/util"
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
		slug, err := util.GeneratePrefixedID(util.Slugify(u.Body.Name), "-", 6)

		if err != nil {
			return nil, err
		}

		endpoint.Slug = slug
		endpoint.Name = u.Body.Name
	}

	if u.Body.Description != "" {
		endpoint.Description = null.NewString(u.Body.Description, true)
	}

	err = u.EndpointRepo.UpdateEndpoint(ctx, endpoint)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}
