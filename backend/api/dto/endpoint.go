package dto

type CreateEndpointRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ProjectID   string `json:"project_id"`
}
