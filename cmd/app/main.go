package main

import (
	"github.com/vicdevcode/schedule/internal/app"
	"github.com/vicdevcode/schedule/pkg/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
