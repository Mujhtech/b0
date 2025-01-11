package services

import (
	"context"

	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/errors"
)

type FindProjectService struct {
	ProjectRepo store.ProjectRepository
	User        *models.User
	ProjectID   string
}

func (f *FindProjectService) Run(ctx context.Context) (*models.Project, error) {

	project, err := f.ProjectRepo.FindProjectByID(ctx, f.ProjectID)

	if err != nil {
		return nil, err
	}

	if project.OwnerID != f.User.ID {
		return nil, errors.ErrNotAuthorized
	}

	return project, nil
}
