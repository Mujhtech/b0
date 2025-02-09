package store

import "github.com/mujhtech/b0/database"

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
