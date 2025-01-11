package dto

type CreateProjectRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
