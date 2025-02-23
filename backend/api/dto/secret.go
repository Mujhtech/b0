package dto

type Secret struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Notes string `json:"notes,omitempty"`
}

type SecretRequestDto struct {
	Secrets []Secret `json:"secrets"`
}
