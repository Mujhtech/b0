package store

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/models"
)

const (
	aiTokenCreditBaseTable    = "ai_token_credits"                                                                                        // #nosec G101
	aiTokenCreditSelectColumn = "id, owner_id, model, credits, total_credits, used_credits, metadata, created_at, updated_at, deleted_at" // #nosec G101
)

type aiTokenCreditRepo struct {
	db *database.Database
}

func NewAITokenCreditRepository(db *database.Database) AITokenCreditRepository {
	return &aiTokenCreditRepo{
		db: db,
	}
}

// FindLogsByProjectID implements AITokenCreditRepository.
func (a *aiTokenCreditRepo) FindAITokenCreditByOwnerID(ctx context.Context, ownerId string) ([]*models.AITokenCredit, error) {
	stmt := Builder.
		Select(aiTokenCreditSelectColumn).
		From(aiTokenCreditBaseTable).
		Where(squirrel.Eq{"owner_id": ownerId}).
		Where(excludeDeleted).
		OrderBy(orderByCreatedAtDesc)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.AITokenCredit{}
	if err := a.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find ai token credits by owner id")
	}

	return dst, nil
}

// DeleteAITokenCredit implements AITokenCreditRepository.
func (a *aiUsageRepo) DeleteAITokenCredit(ctx context.Context, id string) error {
	stmt := Builder.
		Update(aiTokenCreditBaseTable).
		Set("deleted_at", "NOW()").
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = a.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to delete ai token credit")
	}

	return nil
}
