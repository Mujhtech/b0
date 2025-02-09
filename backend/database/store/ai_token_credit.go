package store

import "github.com/mujhtech/b0/database"

const (
	aiTokenCreditBaseTable    = "ai_token_credits"
	aiTokenCreditSelectColumn = "id, owner_id, model, credits, total_credits, used_credits, metadata, created_at, updated_at, deleted_at"
)

type aiTokenCreditRepo struct {
	db *database.Database
}

func NewAITokenCreditRepository(db *database.Database) AITokenCreditRepository {
	return &aiTokenCreditRepo{
		db: db,
	}
}
