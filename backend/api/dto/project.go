package dto

type CreateProjectRequestDto struct {
	Prompt       string `json:"prompt"`
	IsTemplate   bool   `json:"is_template,omitempty"`
	Model        string `json:"model,omitempty"`
	FramekworkID string `json:"framework_id,omitempty"`
}

type ProjectActionRequestDto struct {
	Action string `json:"action"`
}
