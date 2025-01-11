package services

import (
	"context"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/errors"
)

type UpdateProjectService struct {
	Body        *dto.CreateProjectRequestDto
	ProjectRepo store.ProjectRepository
	User        *models.User
	ProjectID   string
}

func (u *UpdateProjectService) Run(ctx context.Context) (*models.Project, error) {

	project, err := u.ProjectRepo.FindProjectByID(ctx, u.ProjectID)

	if err != nil {
		return nil, err
	}

	if project.OwnerID != u.User.ID {
		return nil, errors.ErrNotAuthorized
	}

	if u.Body.Name != "" {
		project.Name = u.Body.Name
		project.Slug = slugify(u.Body.Name)
	}

	if u.Body.Description != "" {
		project.Description = null.NewString(u.Body.Description, true)
	}

	err = u.ProjectRepo.UpdateProject(ctx, project)

	if err != nil {
		return nil, err
	}

	return project, nil
}
