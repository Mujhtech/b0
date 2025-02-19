package store

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/models"
)

type TotalAIUsageFilterRange string

const (
	aiUsageBaseTable    = "ai_usages"
	aiUsageSelectColumn = "id, owner_id, project_id, endpoint_id, input_tokens, output_tokens, model, usage_type, is_premium, metadata, created_at, updated_at, deleted_at"

	TotalAIUsageFilterRangeToday  TotalAIUsageFilterRange = "today"
	TotalAIUsageFilterRangeMonth  TotalAIUsageFilterRange = "month"
	TotalAIUsageFilterRangeCustom TotalAIUsageFilterRange = "custom"
)

type TotalAIUsageFilter struct {
	OwnerID    string                  `json:"owner_id"`
	Model      string                  `json:"model"`
	ProjectID  string                  `json:"project_id"`
	EndpointID string                  `json:"endpoint_id"`
	From       string                  `json:"from"`
	To         string                  `json:"to"`
	Range      TotalAIUsageFilterRange `json:"range"`
	IsPremium  *bool                   `json:"is_premium"`
}

type TotalAIUsage struct {
	TotalUsage             int `db:"total_usage" json:"total_usage"`
	TotalInputTokensMonth  int `db:"total_input_tokens,omitempty" json:"total_input_tokens,omitempty"`
	TotalOutputTokensMonth int `db:"total_output_tokens,omitempty" json:"total_output_tokens,omitempty"`
}

type aiUsageRepo struct {
	db *database.Database
}

func NewAIUsageRepository(db *database.Database) AIUsageRepository {
	return &aiUsageRepo{
		db: db,
	}
}

// CreateAIUsage implements AIUsageRepository.
func (a *aiUsageRepo) CreateAIUsage(ctx context.Context, aiUsage *models.AIUsage) error {
	metadata := "{}"

	if aiUsage.Metadata != nil {
		metadataByte, err := json.Marshal(aiUsage.Metadata)

		if err != nil {
			return err
		}

		metadata = string(metadataByte)
	}

	stmt := Builder.
		Insert(aiUsageBaseTable).
		Columns(
			"id",
			"owner_id",
			"project_id",
			"endpoint_id",
			"input_tokens",
			"output_tokens",
			"model",
			"usage_type",
			"is_premium",
			"metadata",
		).
		Values(
			aiUsage.ID,
			aiUsage.OwnerID,
			aiUsage.ProjectID,
			aiUsage.EndpointID,
			aiUsage.InputToken,
			aiUsage.OutputToken,
			aiUsage.Model,
			aiUsage.UsageType,
			aiUsage.IsPremium,
			metadata,
		)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = a.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to create ai usage")
	}

	return nil
}

// UpdateAIUsage implements ProjectRepository.
func (a *aiUsageRepo) UpdateAIUsage(ctx context.Context, aiUsage *models.AIUsage) error {

	stmt := Builder.
		Update(aiUsageBaseTable).
		Where(squirrel.Eq{"id": aiUsage.ID}).
		Where(excludeDeleted)

	if aiUsage.InputToken != "" {
		stmt = stmt.Set("input_tokens", aiUsage.InputToken)
	}

	if aiUsage.OutputToken != "" {
		stmt = stmt.Set("output_tokens", aiUsage.OutputToken)
	}

	if aiUsage.Model != "" {
		stmt = stmt.Set("model", aiUsage.Model)
	}

	if aiUsage.UsageType != "" {
		stmt = stmt.Set("usage_type", aiUsage.UsageType)
	}

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = a.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to update ai usage")
	}

	return nil
}

// DeleteAIUsage implements AIUsageRepository.
func (a *aiUsageRepo) DeleteAIUsage(ctx context.Context, id string) error {
	stmt := Builder.
		Update(aiUsageBaseTable).
		Set("deleted_at", "NOW()").
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = a.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to delete ai usage")
	}

	return nil
}

// FindAIUsageByID implements AIUsageRepository.
func (a *aiUsageRepo) FindAIUsageByID(ctx context.Context, id string) (*models.AIUsage, error) {
	stmt := Builder.
		Select(aiUsageSelectColumn).
		From(aiUsageBaseTable).
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := new(models.AIUsage)
	if err := a.db.GetDB().GetContext(ctx, dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find ai usage by slug")
	}

	return dst, nil
}

// FindAIUsageByProjectID implements AIUsageRepository.
func (a *aiUsageRepo) FindAIUsageByProjectID(ctx context.Context, projectId string) ([]*models.AIUsage, error) {
	stmt := Builder.
		Select(aiUsageSelectColumn).
		From(aiUsageBaseTable).
		Where(squirrel.Eq{"project_id": projectId}).
		Where(excludeDeleted).
		OrderBy(orderByCreatedAtDesc)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.AIUsage{}
	if err := a.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find ai usages by project id")
	}

	return dst, nil
}

// GetTotalUsage implements AIUsageRepository.
func (a *aiUsageRepo) GetTotalUsage(ctx context.Context, opts TotalAIUsageFilter) (*TotalAIUsage, error) {
	stmt := Builder.
		Select("COUNT(*) AS total_usage").
		From(aiUsageBaseTable).
		Where(excludeDeleted)

	if opts.OwnerID != "" {
		stmt = stmt.Where(squirrel.Eq{"owner_id": opts.OwnerID})
	}

	if opts.Model != "" {
		stmt = stmt.Where(squirrel.Eq{"model": opts.Model})
	}

	if opts.ProjectID != "" {
		stmt = stmt.Where(squirrel.Eq{"project_id": opts.ProjectID})
	}

	if opts.EndpointID != "" {
		stmt = stmt.Where(squirrel.Eq{"endpoint_id": opts.EndpointID})
	}

	if opts.IsPremium != nil {
		stmt = stmt.Where(squirrel.Eq{"is_premium": opts.IsPremium})
	}

	switch opts.Range {
	case TotalAIUsageFilterRangeToday:
		stmt = stmt.Where("DATE(created_at) = CURRENT_DATE")
	case TotalAIUsageFilterRangeMonth:
		stmt = stmt.Where("DATE(created_at) >= DATE_TRUNC('month', CURRENT_DATE)")

	}

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := new(TotalAIUsage)
	if err := a.db.GetDB().GetContext(ctx, dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to get total usage in current month by owner id")
	}

	return dst, nil
}
