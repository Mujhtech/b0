package store

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/models"
)

const (
	endpointBaseTable    = "endpoints"                                                                          // #nosec G101
	endpointSelectColumn = "id, owner_id, project_id, name, slug, metadata, created_at, updated_at, deleted_at" // #nosec G101
)

type endpointRepo struct {
	db *database.Database
}

func NewEndpointRepository(db *database.Database) EndpointRepository {
	return &endpointRepo{
		db: db,
	}
}

// CreateEndpoint implements EndpointRepository.
func (e *endpointRepo) CreateEndpoint(ctx context.Context, endpoint *models.Endpoint) error {
	metadata := "{}"

	if endpoint.Metadata != nil {
		metadataByte, err := json.Marshal(endpoint.Metadata)

		if err != nil {
			return err
		}

		metadata = string(metadataByte)
	}

	stmt := Builder.
		Insert(endpointBaseTable).
		Columns(
			"owner_id",
			"project_id",
			"name",
			"slug",
			"metadata",
		).
		Values(
			endpoint.OwnerID,
			endpoint.ProjectID,
			endpoint.Name,
			endpoint.Slug,
			metadata,
		)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = e.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to create endpoint")
	}

	return nil
}

// FindEndpointByID implements EndpointRepository.
func (e *endpointRepo) FindEndpointByID(ctx context.Context, id string) (*models.Endpoint, error) {
	stmt := Builder.
		Select(endpointSelectColumn).
		From(endpointBaseTable).
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := new(models.Endpoint)
	if err := e.db.GetDB().GetContext(ctx, dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find endpoint by id")
	}

	return dst, nil
}

// FindEndpointByOwnerID implements EndpointRepository.
func (e *endpointRepo) FindEndpointByOwnerID(ctx context.Context, ownerID string) ([]*models.Endpoint, error) {
	stmt := Builder.
		Select(endpointSelectColumn).
		From(endpointBaseTable).
		Where(squirrel.Eq{"owner_id": ownerID}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.Endpoint{}
	if err := e.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find endpoints by owner id")
	}

	return dst, nil
}

// FindEndpointByOwnerID implements EndpointRepository.
func (e *endpointRepo) FindEndpointByProjectID(ctx context.Context, projectID string) ([]*models.Endpoint, error) {
	stmt := Builder.
		Select(endpointSelectColumn).
		From(endpointBaseTable).
		Where(squirrel.Eq{"project_id": projectID}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.Endpoint{}
	if err := e.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find endpoints by project id")
	}

	return dst, nil
}

// UpdateEndpoint implements EndpointRepository.
func (e *endpointRepo) UpdateEndpoint(ctx context.Context, endpoint *models.Endpoint) error {

	stmt := Builder.
		Update(endpointBaseTable).
		Where(squirrel.Eq{"id": endpoint.ID}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = e.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to update endpoint")
	}

	return nil
}

// DeleteEndpoint implements EndpointRepository.
func (e *endpointRepo) DeleteEndpoint(ctx context.Context, id string) error {
	stmt := Builder.
		Update(endpointBaseTable).
		Set("deleted_at", "NOW()").
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = e.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to delete endpoint")
	}

	return nil
}
