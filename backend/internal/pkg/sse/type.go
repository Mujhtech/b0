package sse

import (
	"encoding/json"

	"github.com/mujhtech/b0/internal/pkg/agent"
)

type EventType string

const (
	EventTypeTaskStarted   EventType = "task_started"
	EventTypeTaskUpdate    EventType = "task_updated"
	EventTypeTaskCompleted EventType = "task_completed"
	EventTypeTaskFailed    EventType = "task_failed"
)

type UploadProgressStatus string

const (
	UploadProgressStatusStarted   UploadProgressStatus = "started"
	UploadProgressStatusPending   UploadProgressStatus = "pending"
	UploadProgressStatusUploading UploadProgressStatus = "uploading"
	UploadProgressStatusFailed    UploadProgressStatus = "failed"
	UploadProgressStatusCompleted UploadProgressStatus = "completed"
	UploadProgressStatusCancelled UploadProgressStatus = "cancelled"
)

type AgentData struct {
	Message   string            `json:"message"`
	Error     string            `json:"error,omitempty"`
	Workflows *[]agent.Workflow `json:"workflows,omitempty"`
}

type Event struct {
	Type EventType       `json:"type"`
	Data json.RawMessage `json:"data"`
}
