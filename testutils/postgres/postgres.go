package postgres

import (
	"context"
	"log"
	"runtime/debug"
	"testing"

	"ikea/config"
	"ikea/migrations"
	"ikea/models"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
)

var tables = []interface{}{
	&models.User{},
}

type common interface {
	Run(ctx context.Context) error
	DB() *bun.DB
	WaitConnected(f func(db *bun.DB) error) error
	Cleanup()
}

func Setup(migrationsDir string) (*bun.DB, func()) {
	cfg := config.NewDB()
	instance := NewTestDocker(cfg)

	ctx := context.Background()
	err := instance.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := instance.DB()

	err = instance.WaitConnected(func(db *bun.DB) error {
		_, err := db.Exec("SELECT 1")
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	err = migrations.Run(ctx, db, config.NewConfig())
	if err != nil {
		instance.Cleanup()
		log.Fatal(errors.Wrap(err, "failed to run migrations"))
	}

	return db, instance.Cleanup
}

func truncateTables(t *testing.T, ctx context.Context, db *bun.DB) {
	for _, m := range tables {
		_, err := db.NewTruncateTable().Model(m).Cascade().Exec(ctx)
		require.NoError(t, err)
	}
}

func CleanupAndRecover(t *testing.T, db *bun.DB) {
	truncateTables(t, context.Background(), db)
	if err := recover(); err != nil {
		t.Errorf("panic: %v\n %v", err, string(debug.Stack()))
	}
}
