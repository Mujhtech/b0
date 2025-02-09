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

	project := &models.Project{
		ID:          uuid.New().String(),
		OwnerID:     c.User.ID,
		Name:        c.ProjectTitleAndSlug.Title,
		Model:       null.NewString(c.Body.Model, c.Body.Model != ""),
		Slug:        util.ToLower(util.Slugify(c.ProjectTitleAndSlug.Slug)),
		Description: null.NewString(c.ProjectTitleAndSlug.Description, true),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := c.ProjectRepo.CreateProject(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}
