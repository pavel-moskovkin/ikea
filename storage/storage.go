package storage

import (
	"database/sql"
	"time"

	"localshop/config"
	"localshop/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Storage struct {
	*bun.DB

	config config.Config
}

func NewDB(cfg config.Config) (*Storage, error) {
	// postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	// dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%v",
	// 	cfg.Username, cfg.Password, cfg.Address, cfg.DBName, cfg.SslSecure)

	connector := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(cfg.DbConfig.Address),
		pgdriver.WithInsecure(cfg.DbConfig.Insecure),
		pgdriver.WithUser(cfg.DbConfig.Username),
		pgdriver.WithPassword(cfg.DbConfig.Password),
		pgdriver.WithDatabase(cfg.DbConfig.DBName),
		pgdriver.WithTimeout(5*time.Second), // add these params to config, in case if needed
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqlDB := sql.OpenDB(connector)
	db := bun.NewDB(sqlDB, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	err := db.Ping()
	if err != nil {
		panic("db ping failed")
	}

	return &Storage{
		DB:     db,
		config: cfg,
	}, nil
}

func (s *Storage) Db() *bun.DB {
	return s.DB
}

type UserStore interface {
	UserGet(uuid uuid.UUID) (models.User, error)
	UserCreate(user models.User) (uuid.UUID, error)
}

type UserStorage struct {
	db Storage
}

func (s *UserStorage) UserGet(uuid uuid.UUID) (models.User, error) {
	_, err := s.db.Exec("SELECT 1")
	if err != nil {
		return models.User{}, err
	}
	return models.User{}, nil
}

func (s *UserStorage) UserCreate(user models.User) (uuid.UUID, error) {
	return uuid.Nil, nil
}
