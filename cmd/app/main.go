package main

import (
	"github.com/vicdevcode/init_template/internal/app"
	"github.com/vicdevcode/init_template/pkg/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
