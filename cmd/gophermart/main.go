package main

import (
	"github.com/Kosvu/gophermart/internal/core/config"
	"github.com/Kosvu/gophermart/internal/core/migrations"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	if err := migrations.StartMigration(cfg.DatabaseURI); err != nil {
		panic(err)
	}
}
