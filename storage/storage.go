package storage

import (
	"context"
	"database/sql"
	"errors"

	"ikea/config"
	"ikea/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Storage struct {
	db *bun.DB

	config config.Config
}

func NewDB(cfg config.Config) (*Storage, error) {
	// postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	// dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%v",
	// 	cfg.Username, cfg.Password, cfg.Address, cfg.DBName, cfg.Insecure)

	connector := NewConnectorFromConfig(cfg.DB)

	sqlDB := sql.OpenDB(connector)
	db := bun.NewDB(sqlDB, pgdialect.New())

	db.AddQueryHook(CustomQueryHook())

	err := db.Ping()
	if err != nil {
		panic("db ping failed")
	}

	return &Storage{
		db:     db,
		config: cfg,
	}, nil
}

func (s *Storage) Db() *bun.DB {
	return s.db
}

type DataStore interface {
	UserCreate(user models.User) (uuid.UUID, error)
	UserGet(uuid uuid.UUID) (models.User, error)
	UserUpdate(user models.User) error
	UserDelete(uuid uuid.UUID) error

	OrderCreate(order models.Order) (int, error)
	OrderGet(number int) (models.Order, error)
	OrderUpdate(order models.Order) error
	OrderDelete(number int) error

	ItemCreate(item models.Item) (uuid.UUID, error)
	ItemGet(uuid uuid.UUID) (models.Item, error)
	ItemUpdate(item models.Item) error
	ItemDelete(uuid uuid.UUID) error
}

func NewStorageFromDB(db *bun.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) UserGet(ctx context.Context, uuid uuid.UUID) (models.User, error) {
	user := models.User{UUID: &uuid}
	err := s.db.NewSelect().Model(&user).WherePK().Scan(ctx)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Storage) UserCreate(ctx context.Context, user models.User) (uuid.UUID, error) {
	res, err := s.db.NewInsert().Model(&user).Returning("uuid").Exec(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return uuid.Nil, errors.New("invalid rows affected count")
	}

	return *user.UUID, nil
}
