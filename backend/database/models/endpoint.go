package models

import (
	"time"

	"github.com/guregu/null"
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

type Endpoint struct {
	ID          string         `json:"id" db:"id"`
	OwnerID     string         `json:"owner_id" db:"owner_id"`
	ProjectID   string         `json:"project_id" db:"project_id"`
	Name        string         `json:"name" db:"name"`
	Path        string         `json:"path" db:"path"`
	Method      EndpointMethod `json:"method" db:"method"`
	Description null.String    `json:"description" db:"description"`
	IsPublic    bool           `json:"is_public" db:"is_public"`
	Status      EndpointStatus `json:"status" db:"status"`
	Metadata    interface{}    `json:"metadata" db:"metadata"`
	CreatedAt   time.Time      `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt   null.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}
