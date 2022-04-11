package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"ikea/config"

	_ "github.com/lib/pq"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type dockerTest struct {
	db  *bun.DB
	cfg config.DB

	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func NewTestDocker(cfg config.DB) *dockerTest {
	return &dockerTest{
		cfg: cfg,
	}
}

func (d *dockerTest) Run(ctx context.Context) error {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return errors.Wrap(err, "could not connect to docker")
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_DB=" + d.cfg.DBName,
			"POSTGRES_USER=" + d.cfg.Username,
			"POSTGRES_PASSWORD=" + d.cfg.Password,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return errors.Wrap(err, "could not start resource")
	}

	d.pool = pool
	d.resource = resource

	hostAndPort := resource.GetHostPort("5432/tcp")
	d.cfg.Address = hostAndPort

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", d.cfg.Username, d.cfg.Password, hostAndPort, d.cfg.DBName)
	log.Println("Connecting to database on url: ", databaseUrl)
	// _ = resource.Expire(120)

	// var db *sql.DB
	// pool.MaxWait = 10 * time.Second
	// if err = pool.Retry(func() error {
	// 	db, err = sql.Open("postgres", databaseUrl)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return db.Ping()
	// }); err != nil {
	// 	log.Fatalf("Could not connect to docker: %s", err)
	// }

	// TODO SetTimeZoneUTC

	connector := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(d.cfg.Address),
		pgdriver.WithUser(d.cfg.Username),
		pgdriver.WithPassword(d.cfg.Password),
		pgdriver.WithDatabase(d.cfg.DBName),
		pgdriver.WithInsecure(d.cfg.Insecure),
		pgdriver.WithTimeout(5*time.Second), // add these params to config, in case if needed
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqlDB := sql.OpenDB(connector)
	bundb := bun.NewDB(sqlDB, pgdialect.New())

	d.db = bundb

	return nil
}

func (d dockerTest) DB() *bun.DB {
	return d.db
}

func (d dockerTest) WaitConnected(f func(db *bun.DB) error) error {
	return d.pool.Retry(func() error { return f(d.db) })
}

func (d *dockerTest) Cleanup() {
	if d.db != nil {
		err := d.db.Close()
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to close db connection"))
		}
	}
	d.db = nil

	// When you're done, kill and remove the container
	if d.pool != nil {
		fmt.Printf("removing container")
		err := d.resource.Close()
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to close db connection"))
		}
	}
	d.pool = nil
}
