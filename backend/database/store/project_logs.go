package store

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/models"
)

const (
	projectLogBaseTable    = "project_logs"
	projectLogSelectColumn = "id, owner_id, project_id, endpoint_id, log_type, log_data, metadata, created_at, updated_at, deleted_at"
)

type projectLogRepo struct {
	db *database.Database
}

func NewProjectLogRepository(db *database.Database) ProjectLogRepository {
	return &projectLogRepo{
		db: db,
	}
}

// FindLogsByProjectID implements ProjectLogRepository.
func (p *projectLogRepo) FindLogsByProjectID(ctx context.Context, projectId string) ([]*models.ProjectLog, error) {
	stmt := Builder.
		Select(projectLogSelectColumn).
		From(projectLogBaseTable).
		Where(squirrel.Eq{"project_id": projectId}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.ProjectLog{}
	if err := p.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find project logs by project id")
	}

	return dst, nil
}

// DeleteAIUsage implements ProjectLogRepository.
func (p *projectLogRepo) DeleteLog(ctx context.Context, id string) error {
	stmt := Builder.
		Update(projectLogBaseTable).
		Set("deleted_at", "NOW()").
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to delete project log")
	}

	return nil
}
