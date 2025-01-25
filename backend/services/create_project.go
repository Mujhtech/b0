package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/util"
)

type CreateProjectService struct {
	ProjectTitleAndSlug *agent.ProjectTitleAndSlug
	Body                *dto.CreateProjectRequestDto
	ProjectRepo         store.ProjectRepository
	User                *models.User
}

func (c *CreateProjectService) Run(ctx context.Context) (*models.Project, error) {

	slug, err := util.GeneratePrefixedID(util.Slugify(c.ProjectTitleAndSlug.Slug), "-", 6)

	if err != nil {
		return nil, err
	}

	project := &models.Project{
		ID:      uuid.New().String(),
		OwnerID: c.User.ID,
		Name:    c.ProjectTitleAndSlug.Title,
		Slug:    util.ToLower(slug),
		//Description: null.NewString(c.Body.Description, c.Body.Description != ""),
		Metadata:  null.NewString(c.ProjectTitleAndSlug.Description, true),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = c.ProjectRepo.CreateProject(ctx, project)

	if err != nil {
		return nil, err
	}

	return project, nil
}
