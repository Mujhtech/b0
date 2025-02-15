package models

import (
	"time"

	"github.com/guregu/null"
)

type Project struct {
	ID          string      `json:"id" db:"id"`
	OwnerID     string      `json:"owner_id" db:"owner_id"`
	Name        string      `json:"name" db:"name"`
	Slug        string      `json:"slug" db:"slug"`
	Description null.String `json:"description" db:"description"`
	Model       null.String `json:"model" db:"model"`
	ContainerID null.String `json:"-" db:"container_id"`
	Port        null.String `json:"port" db:"port"`
	ServerUrl   null.String `json:"server_url" db:"server_url"`
	Framework   string      `json:"framework" db:"framework"`
	Language    string      `json:"language" db:"language"`
	Metadata    interface{} `json:"metadata" db:"metadata"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt   null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
