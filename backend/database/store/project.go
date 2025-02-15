package store

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/models"
)

const (
	projectBaseTable    = "projects"
	projectSelectColumn = "id, owner_id, name, slug, description, model, server_url, port, framework, language, metadata, created_at, updated_at, deleted_at"
)

type projectRepo struct {
	db *database.Database
}

func NewProjectRepository(db *database.Database) ProjectRepository {
	return &projectRepo{
		db: db,
	}
}

// CreateProject implements ProjectRepository.
func (p *projectRepo) CreateProject(ctx context.Context, project *models.Project) error {
	metadata := "{}"

	if project.Metadata != nil {
		metadataByte, err := json.Marshal(project.Metadata)

		if err != nil {
			return err
		}

		metadata = string(metadataByte)
	}

	stmt := Builder.
		Insert(projectBaseTable).
		Columns(
			"id",
			"owner_id",
			"name",
			"slug",
			"description",
			"model",
			"container_id",
			"framework",
			"language",
			"port",
			"server_url",
			"metadata",
		).
		Values(
			project.ID,
			project.OwnerID,
			project.Name,
			project.Slug,
			project.Description,
			project.Model,
			project.ContainerID,
			project.Framework,
			project.Language,
			project.Port,
			project.ServerUrl,
			metadata,
		)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to create project")
	}

	return nil
}

// FindProjectByID implements ProjectRepository.
func (p *projectRepo) FindProjectByID(ctx context.Context, id string) (*models.Project, error) {
	stmt := Builder.
		Select(projectSelectColumn).
		From(projectBaseTable).
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := new(models.Project)
	if err := p.db.GetDB().GetContext(ctx, dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find project by id")
	}

	return dst, nil
}

// FindProjectBySlug implements ProjectRepository.
func (p *projectRepo) FindProjectBySlug(ctx context.Context, slug string) (*models.Project, error) {
	stmt := Builder.
		Select(projectSelectColumn).
		From(projectBaseTable).
		Where(squirrel.Eq{"slug": slug}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := new(models.Project)
	if err := p.db.GetDB().GetContext(ctx, dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find project by slug")
	}

	return dst, nil
}

// FindProjectByOwnerID implements ProjectRepository.
func (p *projectRepo) FindProjectByOwnerID(ctx context.Context, ownerID string) ([]*models.Project, error) {
	stmt := Builder.
		Select(projectSelectColumn).
		From(projectBaseTable).
		Where(squirrel.Eq{"owner_id": ownerID}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return nil, err
	}

	dst := []*models.Project{}
	if err := p.db.GetDB().SelectContext(ctx, &dst, sql, args...); err != nil {
		return nil, ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to find projects by owner id")
	}

	return dst, nil
}

// UpdateProject implements ProjectRepository.
func (p *projectRepo) UpdateProject(ctx context.Context, project *models.Project) error {

	stmt := Builder.
		Update(projectBaseTable).
		Where(squirrel.Eq{"id": project.ID}).
		Where(excludeDeleted)

	if project.ContainerID.Valid && project.ContainerID.String != "" {
		stmt = stmt.Set("container_id", project.ContainerID)
	}

	if project.Port.Valid && project.Port.String != "" {
		stmt = stmt.Set("port", project.Port)
	}

	if project.ServerUrl.Valid && project.ServerUrl.String != "" {
		stmt = stmt.Set("server_url", project.ServerUrl)
	}

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to update project")
	}

	return nil
}

// DeleteProject implements ProjectRepository.
func (p *projectRepo) DeleteProject(ctx context.Context, id string) error {
	stmt := Builder.
		Update(projectBaseTable).
		Set("deleted_at", "NOW()").
		Where(squirrel.Eq{"id": id}).
		Where(excludeDeleted)

	sql, args, err := stmt.ToSql()

	if err != nil {
		return err
	}

	_, err = p.db.GetDB().ExecContext(ctx, sql, args...)

	if err != nil {
		return ProcessSQLErrorfWithCtx(ctx, sql, err, "failed to delete project")
	}

	return nil
}
