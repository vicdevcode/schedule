package main

import (
	"log"

	"github.com/vicdevcode/schedule/internal/entity"
	"github.com/vicdevcode/schedule/pkg/config"
	"github.com/vicdevcode/schedule/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New((*postgres.Config)(&cfg.DB))
	if err != nil {
		log.Fatalf("Failed connect to postgres: %s", err.Error())
		return
	}
	db.AutoMigrate(&entity.WeeklyEvent{})
}
