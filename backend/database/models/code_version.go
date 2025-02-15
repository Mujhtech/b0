package models

import (
	"time"

	"github.com/guregu/null"
)

type CodeVersion struct {
	ID         string      `json:"id" db:"id"`
	OwnerID    string      `json:"owner_id" db:"owner_id"`
	ProjectID  string      `json:"project_id" db:"project_id"`
	EndpointID null.String `json:"endpoint_id" db:"endpoint_id"`
	Version    string      `json:"version" db:"version"`
	CommitID   string      `json:"commit_id" db:"commit_id"`
	Branch     string      `json:"branch" db:"branch"`
	CommitMsg  string      `json:"commit_msg" db:"commit_msg"`
	Content    interface{} `json:"content" db:"content"`
	Metadata   interface{} `json:"metadata" db:"metadata"`
	CreatedAt  time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt  time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt  null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
