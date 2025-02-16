package dto

import (
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/internal/pkg/agent"
)

type CreateEndpointRequestDto struct {
	Name        string                `json:"name"`
	Description string                `json:"description,omitempty"`
	ProjectID   string                `json:"project_id"`
	Path        string                `json:"path"`
	Method      models.EndpointMethod `json:"method"`
	IsPublic    bool                  `json:"is_public"`
	Status      models.EndpointStatus `json:"status"`
}

type GetEndpointQuery struct {
	ProjectID string `json:"project_id,omitempty"`
}

type UpdateEndpointWorkflowRequestDto struct {
	Workflows []*agent.Workflow `json:"workflows"`
}
