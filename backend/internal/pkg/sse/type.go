package sse

import (
	"encoding/json"
)

type EventType string

var (
	FailedToPublishTaskUpdatedEvent   = "failed to publish task updated event"
	FailedToPublishTaskCompletedEvent = "failed to publish task completed event"
	FailedToPublishTaskFailedEvent    = "failed to publish task failed event"
	FailedToPublishTaskStartedEvent   = "failed to publish task started event"
)

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

type Event struct {
	Type EventType       `json:"type"`
	Data json.RawMessage `json:"data"`
}
