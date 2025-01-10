package services

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
)

type CreateProjectService struct {
	Body        *dto.CreateProjectRequestDto
	ProjectRepo store.ProjectRepository
	User        *models.User
}

func (c *CreateProjectService) Run(ctx context.Context) (*models.Project, error) {

	slug := slugify(c.Body.Name)

	project := &models.Project{
		ID:          uuid.New().String(),
		OwnerID:     c.User.ID,
		Name:        c.Body.Name,
		Slug:        slug,
		Description: null.NewString(c.Body.Description, c.Body.Description != ""),
		Metadata:    null.NewString("{}", true),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := c.ProjectRepo.CreateProject(ctx, project)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func slugify(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
