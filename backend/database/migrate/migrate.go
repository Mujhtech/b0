package migrate

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/jmoiron/sqlx"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database"
	"maragu.dev/migrate"
)

//go:embed postgres/*.sql
var postgres embed.FS

const (
	postgresSourceDir = "postgres"
	sqliteSourceDir   = "sqlite"
)

func Migrator(ctx context.Context, cfg *config.Config, db *database.Database) (*migrate.Migrator, error) {
	opts, err := getMigratorOpt(cfg.Database.Driver, db.GetDB())
	if err != nil {
		return nil, fmt.Errorf("failed to get migrator opt: %w", err)
	}
	return migrate.New(opts), nil
}

func getMigratorOpt(dbDriver config.DatabaseDriver, db *sqlx.DB) (migrate.Options, error) {

	opts := migrate.Options{
		FS: postgres,
		DB: db.DB,
	}

	folder, _ := fs.Sub(postgres, postgresSourceDir)

	opts.FS = folder

	return opts, nil
}
