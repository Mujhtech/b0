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

		slug, err := util.GeneratePrefixedID(util.Slugify(u.Body.Name), "-", 6)

		if err != nil {
			return nil, err
		}

		project.Slug = slug
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
