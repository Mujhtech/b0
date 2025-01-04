package auth

import (
	"github.com/mujhtech/b0/database/models"
)

type UserSession struct {
	User     *models.User
	Metadata *TokenMetadata
}

type TokenMetadata struct {
	Type     CredentialType
	Metadata interface{}
	TokenID  string
}
