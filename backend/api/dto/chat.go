package dto

type ChatRequestDto struct {
	Text  string `json:"text"`
	Model string `json:"model,omitempty"`
}
