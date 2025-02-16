package store

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/internal/util"
)

const (
	endpointBaseTable    = "endpoints"
	endpointSelectColumn = "id, owner_id, project_id, name, description, path, method, is_public, connectors, workflows, code_generation, status, metadata, created_at, updated_at, deleted_at"
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
	workflows := "{}"
	connectors := "{}"
	codeGeneration := "{}"

	if endpoint.Metadata != nil {
		metadataOutput, err := util.MarshalJSONToString(endpoint.Metadata)
		if err != nil {
			return err
		}
		metadata = metadataOutput
	}

	if endpoint.Workflows != nil {
		workflowsOutput, err := util.MarshalJSONToString(endpoint.Workflows)
		if err != nil {
			return err
		}
		workflows = workflowsOutput
	}

	if endpoint.Connectors != nil {
		connectorsOutput, err := util.MarshalJSONToString(endpoint.Connectors)
		if err != nil {
			return err
		}
		connectors = connectorsOutput
	}

	if endpoint.CodeGeneration != nil {
		codeGenerationOutput, err := util.MarshalJSONToString(endpoint.CodeGeneration)
		if err != nil {
			return err
		}
		codeGeneration = codeGenerationOutput
	}

	stmt := Builder.
		Insert(endpointBaseTable).
		Columns(
			"id",
			"owner_id",
			"project_id",
			"name",
			"description",
			"path",
			"method",
			"is_public",
			"status",
			"connectors",
			"workflows",
			"code_generation",
			"metadata",
		).
		Values(
			endpoint.ID,
			endpoint.OwnerID,
			endpoint.ProjectID,
			endpoint.Name,
			endpoint.Description,
			endpoint.Path,
			endpoint.Method,
			endpoint.IsPublic,
			endpoint.Status,
			connectors,
			workflows,
			codeGeneration,
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
		Where(excludeDeleted).
		OrderBy(orderByCreatedAtDesc)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.EndpointFromDB{}
	if err := e.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find endpoints by owner id")
	}

	return models.ToEndpoints(dst), nil
}

// FindEndpointByOwnerID implements EndpointRepository.
func (e *endpointRepo) FindEndpointByProjectID(ctx context.Context, projectID string) ([]*models.Endpoint, error) {
	stmt := Builder.
		Select(endpointSelectColumn).
		From(endpointBaseTable).
		Where(squirrel.Eq{"project_id": projectID}).
		Where(excludeDeleted).
		OrderBy(orderByCreatedAtDesc)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.EndpointFromDB{}
	if err := e.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find endpoints by project id")
	}

	return models.ToEndpoints(dst), nil
}

// UpdateEndpoint implements EndpointRepository.
func (e *endpointRepo) UpdateEndpoint(ctx context.Context, id string, endpoint *models.Endpoint) error {

	stmt := Builder.
		Update(endpointBaseTable).
		// Set("name", endpoint.Name).
		// Set("description", endpoint.Description).
		// Set("path", endpoint.Path).
		// Set("method", endpoint.Method).
		// Set("status", endpoint.Status).
		// Set("metadata", endpoint.Metadata).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	if endpoint.Workflows != nil {

		workflowsOutput, err := util.MarshalJSONToString(endpoint.Workflows)
		if err != nil {
			return err
		}

		stmt = stmt.Set("workflows", workflowsOutput)
	}

	if endpoint.CodeGeneration != nil {

		codeGenerationOutput, err := util.MarshalJSONToString(endpoint.CodeGeneration)
		if err != nil {
			return err
		}

		stmt = stmt.Set("code_generation", codeGenerationOutput)
	}

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
