package dto

type Secret struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Note      string `json:"note,omitempty"`
	Protected bool   `json:"protected,omitempty"`
}

type SecretRequestDto struct {
	Secrets []Secret `json:"secrets"`
}
