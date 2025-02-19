package models

import (
	"time"

	"github.com/guregu/null"
)

type AIUsage struct {
	ID          string      `json:"id" db:"id"`
	OwnerID     string      `json:"owner_id" db:"owner_id"`
	ProjectID   string      `json:"project_id" db:"project_id"`
	EndpointID  null.String `json:"endpoint_id" db:"endpoint_id"`
	InputToken  string      `json:"input_tokens" db:"input_tokens"`
	OutputToken string      `json:"output_tokens" db:"output_tokens"`
	Model       string      `json:"model" db:"model"`
	UsageType   string      `json:"usage_type" db:"usage_type"`
	IsPremium   bool        `json:"is_premium" db:"is_premium"`
	Metadata    interface{} `json:"metadata" db:"metadata"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt   null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
