package models

import (
	"time"

	"github.com/guregu/null"
)

type AITokenCredit struct {
	ID           string      `json:"id" db:"id"`
	OwnerID      string      `json:"owner_id" db:"owner_id"`
	Model        string      `json:"model" db:"model"`
	Credits      int         `json:"credits" db:"credits"`
	TotalCredits int         `json:"total_credits" db:"total_credits"`
	UsedCredits  int         `json:"used_credits" db:"used_credits"`
	Metadata     interface{} `json:"metadata" db:"metadata"`
	CreatedAt    time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt    null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
