package job

import "time"

type JobName string
type QueueName string

const (
	JobNameWebhook        JobName = "webhook"
	JobNameWorkflowCreate JobName = "workflow.create"
	JobNameWorkflowUpdate JobName = "workflow.update"
	JobNameProjectDeploy  JobName = "project.project"
	JobNameProjectExport  JobName = "project.export"

	QueueNameDefault QueueName = "default"
)

type ClientPayload struct {
	Data  []byte        `json:"data"`
	Delay time.Duration `json:"delay"`
}
