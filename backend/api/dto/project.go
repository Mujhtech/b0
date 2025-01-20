package dto

type CreateProjectRequestDto struct {
	Prompt     string `json:"prompt"`
	IsTemplate bool   `json:"is_template,omitempty"`
}
