package main

import (
	"context"
	"fmt"
	"os"

	"localshop/config"
	"localshop/logger"
	"localshop/migrations"
	"localshop/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config file")
	}

	if len(os.Args) > 1 {
		fmt.Printf("%v\n", os.Args) // TODO remove
		switch os.Args[1] {
		case "migrate":
			runMigrations(cfg)
		default:
			fmt.Printf("Usage: run with flag 'migrate' or without.")
			os.Exit(0)
		}
	} else {
		runService(cfg)
	}
}

func runService(cfg config.Config) {
	log := logger.New(cfg.LogLevel)
	log.Info("hello and welcome!")
}

func runMigrations(cfg config.Config) {
	fmt.Printf("Start migrations:")

	store, err := storage.NewDB(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to init database: %s", err.Error()))
	}

	ctx := context.Background()
	err = migrations.Run(ctx, store.DB, cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to run migrations: %s", err.Error()))
	}
}
