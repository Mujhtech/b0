package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/internal/pkg/agent"
)

type EndpointMethod string
type EndpointStatus string

const (
	EndpointMethodGet    EndpointMethod = "GET"
	EndpointMethodPost   EndpointMethod = "POST"
	EndpointMethodPut    EndpointMethod = "PUT"
	EndpointMethodPatch  EndpointMethod = "PATCH"
	EndpointMethodDelete EndpointMethod = "DELETE"

	EndpointStatusActive   EndpointStatus = "active"
	EndpointStatusInactive EndpointStatus = "inactive"
	EndpointStatusDraft    EndpointStatus = "draft"
)

type WorkflowsData []*agent.Workflow

type Endpoint struct {
	ID             string            `json:"id"`
	OwnerID        string            `json:"owner_id"`
	ProjectID      string            `json:"project_id"`
	Name           string            `json:"name"`
	Path           string            `json:"path"`
	Method         EndpointMethod    `json:"method"`
	Description    null.String       `json:"description"`
	IsPublic       bool              `json:"is_public"`
	Status         EndpointStatus    `json:"status"`
	Metadata       interface{}       `json:"metadata"`
	Connectors     interface{}       `json:"connectors,omitempty"`
	Workflows      []*agent.Workflow `json:"workflows,omitempty"`
	CodeGeneration interface{}       `json:"code_generation,omitempty"`
	CreatedAt      time.Time         `json:"created_at,omitempty"`
	UpdatedAt      time.Time         `json:"updated_at,omitempty"`
	DeletedAt      null.Time         `json:"deleted_at,omitempty"`
}

type EndpointFromDB struct {
	ID             string         `db:"id"`
	OwnerID        string         `db:"owner_id"`
	ProjectID      string         `db:"project_id"`
	Name           string         `db:"name"`
	Path           string         `db:"path"`
	Method         EndpointMethod `db:"method"`
	Description    null.String    `db:"description"`
	IsPublic       bool           `db:"is_public"`
	Status         EndpointStatus `db:"status"`
	Metadata       JSONField      `db:"metadata"`
	Connectors     JSONField      `db:"connectors"`
	Workflows      JSONField      `db:"workflows"`
	CodeGeneration JSONField      `db:"code_generation"`
	CreatedAt      time.Time      `db:"created_at,omitempty"`
	UpdatedAt      time.Time      `db:"updated_at,omitempty"`
	DeletedAt      null.Time      `db:"deleted_at"`
}

func (w *WorkflowsData) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", value)
	}

	return json.Unmarshal(bytes, w)
}

func (e *EndpointFromDB) UnmarshalWorkflows(target interface{}) error {
	return json.Unmarshal(e.Workflows, target)
}

func (e *EndpointFromDB) UnmarshalConnectors(target interface{}) error {
	return json.Unmarshal(e.Connectors, target)
}

func (e *EndpointFromDB) UnmarshalMetadata(target interface{}) error {
	return json.Unmarshal(e.Metadata, target)
}

func (e *EndpointFromDB) UnmarshalCodeGeneration(target interface{}) error {
	return json.Unmarshal(e.CodeGeneration, target)
}

func ToEndpoint(e *EndpointFromDB) *Endpoint {

	var workflows WorkflowsData
	var metadata interface{}
	var connectors interface{}
	var codeGeneration interface{}

	e.UnmarshalWorkflows(&workflows)
	e.UnmarshalMetadata(&metadata)
	e.UnmarshalConnectors(&connectors)
	e.UnmarshalCodeGeneration(&codeGeneration)

	return &Endpoint{
		ID:             e.ID,
		OwnerID:        e.OwnerID,
		ProjectID:      e.ProjectID,
		Name:           e.Name,
		Path:           e.Path,
		Method:         e.Method,
		Description:    e.Description,
		IsPublic:       e.IsPublic,
		Status:         e.Status,
		Metadata:       metadata,
		Connectors:     connectors,
		Workflows:      workflows,
		CodeGeneration: codeGeneration,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		DeletedAt:      e.DeletedAt,
	}
}

func ToEndpoints(endpoints []*EndpointFromDB) []*Endpoint {
	var result []*Endpoint

	for _, e := range endpoints {
		result = append(result, ToEndpoint(e))
	}

	return result
}
