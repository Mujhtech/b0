package store

import (
	"context"

	"github.com/mujhtech/b0/database/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByID(ctx context.Context, id string) (*models.User, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, token *models.Token) error
	FindTokenByID(ctx context.Context, id string) (*models.Token, error)
	DeleteToken(ctx context.Context, id string) error
}

type EndpointRepository interface {
	CreateEndpoint(ctx context.Context, endpoint *models.Endpoint) error
	UpdateEndpoint(ctx context.Context, endpoint *models.Endpoint) error
	FindEndpointByID(ctx context.Context, id string) (*models.Endpoint, error)
	FindEndpointByProjectID(ctx context.Context, projectID string) ([]*models.Endpoint, error)
	FindEndpointByOwnerID(ctx context.Context, ownerID string) ([]*models.Endpoint, error)
	DeleteEndpoint(ctx context.Context, id string) error
}

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *models.Project) error
	UpdateProject(ctx context.Context, project *models.Project) error
	FindProjectBySlug(ctx context.Context, slug string) (*models.Project, error)
	FindProjectByID(ctx context.Context, id string) (*models.Project, error)
	FindProjectByOwnerID(ctx context.Context, ownerID string) ([]*models.Project, error)
	DeleteProject(ctx context.Context, id string) error
}

type AIUsageRepository interface {
	CreateAIUsage(ctx context.Context, aiUsage *models.AIUsage) error
	UpdateAIUsage(ctx context.Context, aiUsage *models.AIUsage) error
	DeleteAIUsage(ctx context.Context, id string) error
	FindAIUsageByID(ctx context.Context, id string) (*models.AIUsage, error)
	FindAIUsageByProjectID(ctx context.Context, projectID string) ([]*models.AIUsage, error)
	GetTotalUsageInCurrentMonth(ctx context.Context, projectID string) (*TotalAIUsageInCurrentMonth, error)
	GetTotalUsageInCurrentDay(ctx context.Context, projectID string) (*TotalAIUsageInCurrentDay, error)
}

type ProjectLogRepository interface{}

type AITokenCreditRepository interface{}
