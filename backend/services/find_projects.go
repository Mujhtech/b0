package services

import (
	"context"

	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
)

type FindProjectsService struct {
	ProjectRepo store.ProjectRepository
	User        *models.User
}

func (f *FindProjectsService) Run(ctx context.Context) ([]*models.Project, error) {

	projects, err := f.ProjectRepo.FindProjectByOwnerID(ctx, f.User.ID)

	if err != nil {
		return nil, err
	}

	return projects, nil
}
