package handlers

import (
	"context"

	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/rs/zerolog"
)

func sendEvent(ctx context.Context, projectID string, eventType sse.EventType, data AgentData, event sse.Streamer) {

	var errorMsg string

	switch eventType {
	case sse.EventTypeTaskStarted:
		errorMsg = sse.FailedToPublishTaskStartedEvent
	case sse.EventTypeTaskUpdate:
		errorMsg = sse.FailedToPublishTaskUpdatedEvent
	case sse.EventTypeTaskCompleted:
		errorMsg = sse.FailedToPublishTaskCompletedEvent
	case sse.EventTypeTaskFailed:
		errorMsg = sse.FailedToPublishTaskFailedEvent
	default:
		errorMsg = "unknown event type"
	}

	if err := event.Publish(ctx, projectID, eventType, data); err != nil {
		zerolog.Ctx(ctx).Error().Msgf("%s: %v", errorMsg, err)
	}
}
