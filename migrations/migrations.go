package migrations

import (
	"context"

	"ikea/config"
	"ikea/scripts/migrations"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func Run(ctx context.Context, db *bun.DB, config config.Config) error {
	migrator := migrate.NewMigrator(db, migrations.Migrations)
	err := migrator.Init(ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to run init migration")
	}

	_, err = migrator.Migrate(ctx)
	if err != nil {
		return errors.Wrap(err, "Error run migrations")
	}

	return nil
}
