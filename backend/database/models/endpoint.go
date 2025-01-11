package models

import (
	"time"

	"github.com/guregu/null"
)

type Endpoint struct {
	ID          string      `json:"id" db:"id"`
	OwnerID     string      `json:"owner_id" db:"owner_id"`
	ProjectID   string      `json:"project_id" db:"project_id"`
	Name        string      `json:"name" db:"name"`
	Slug        string      `json:"slug" db:"slug"`
	Description null.String `json:"description" db:"description"`
	Metadata    interface{} `json:"metadata" db:"metadata"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt   null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
