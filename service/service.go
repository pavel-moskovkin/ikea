package service

import (
	"ikea/config"
	"ikea/storage"
)

type Service struct {
	DB  *storage.Storage
	cfg config.Config
}

func New(cfg config.Config) (*Service, error) {
	db, err := storage.NewDB(cfg)
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		DB:  db,
		cfg: cfg,
	}, nil
}
