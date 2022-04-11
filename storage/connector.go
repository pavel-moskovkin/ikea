package storage

import (
	"time"

	"ikea/config"

	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewConnectorFromConfig(cfg config.DB) *pgdriver.Connector {
	return pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(cfg.Address),
		pgdriver.WithUser(cfg.Username),
		pgdriver.WithPassword(cfg.Password),
		pgdriver.WithDatabase(cfg.DBName),
		pgdriver.WithInsecure(cfg.Insecure),
		pgdriver.WithTimeout(5*time.Second), // add these params to config, in case if needed
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)
}

func CustomQueryHook() *bundebug.QueryHook {
	return bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	)
}
