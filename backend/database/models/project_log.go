package models

import (
	"time"

	"github.com/guregu/null"
)

type ProjectLog struct {
	ID         string      `json:"id" db:"id"`
	OwnerID    string      `json:"owner_id" db:"owner_id"`
	ProjectID  string      `json:"project_id" db:"project_id"`
	EndpointID null.String `json:"endpoint_id" db:"endpoint_id"`
	LogType    string      `json:"log_type" db:"log_type"`
	LogData    interface{} `json:"log_data" db:"log_data"`
	Metadata   interface{} `json:"metadata" db:"metadata"`
	CreatedAt  time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt  time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt  null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
