package main

import (
	"log"

	"github.com/vicdevcode/init_template/internal/entity"
	"github.com/vicdevcode/init_template/pkg/config"
	"github.com/vicdevcode/init_template/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New((*postgres.Config)(&cfg.DB))
	if err != nil {
		log.Fatalf("Failed connect to postgres: %s", err.Error())
		return
	}
	db.AutoMigrate(&entity.Example{})
}
