package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/errors"
	"github.com/mujhtech/b0/internal/util"
)

type CreateEndpointService struct {
	Body         *dto.CreateEndpointRequestDto
	ProjectRepo  store.ProjectRepository
	EndpointRepo store.EndpointRepository
	User         *models.User
}

func (c *CreateEndpointService) Run(ctx context.Context) (*models.Endpoint, error) {

	project, err := c.ProjectRepo.FindProjectByID(ctx, c.Body.ProjectID)

	if err != nil {
		return nil, err
	}

	if project.OwnerID != c.User.ID {
		return nil, errors.ErrNotAuthorized
	}

	slug, err := util.GeneratePrefixedID(util.Slugify(c.Body.Name), "-", 6)

	if err != nil {
		return nil, err
	}

	endpoint := &models.Endpoint{
		ID:          uuid.New().String(),
		OwnerID:     c.User.ID,
		Name:        c.Body.Name,
		Slug:        slug,
		Description: null.NewString(c.Body.Description, c.Body.Description != ""),
		Metadata:    null.NewString("{}", true),
		ProjectID:   project.OwnerID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = c.EndpointRepo.CreateEndpoint(ctx, endpoint)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}
