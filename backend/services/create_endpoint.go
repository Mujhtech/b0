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

	endpoint := &models.Endpoint{
		ID:          uuid.New().String(),
		OwnerID:     c.User.ID,
		ProjectID:   project.ID,
		Name:        c.Body.Name,
		Description: null.NewString(c.Body.Description, c.Body.Description != ""),
		Path:        c.Body.Path,
		Method:      c.Body.Method,
		Metadata:    null.NewString("{}", true),
		IsPublic:    false,
		Status:      models.EndpointStatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = c.EndpointRepo.CreateEndpoint(ctx, endpoint)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}
