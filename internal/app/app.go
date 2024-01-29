package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	v1 "github.com/vicdevcode/init_template/internal/controller/http/v1"
	"github.com/vicdevcode/init_template/internal/usecase"
	"github.com/vicdevcode/init_template/internal/usecase/repo"
	"github.com/vicdevcode/init_template/pkg/config"
	"github.com/vicdevcode/init_template/pkg/httpserver"
	"github.com/vicdevcode/init_template/pkg/logger"
	"github.com/vicdevcode/init_template/pkg/postgres"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Env)

	log.Info(fmt.Sprintf("Starting server at Port: %s", cfg.Http.Port))

	db, err := postgres.New((*postgres.Config)(&cfg.DB))
	if err != nil {
		logger.Fatal(log, "Failed connect to postgres:", err)
	}
	log.Info("Connected to postgres")

	// UseCases
	exampleUseCase := usecase.NewExample(repo.NewExample(db), cfg.ContextTimeout)

	// HTTP SERVER
	gin.SetMode(gin.ReleaseMode)
	if cfg.Env == "local" {
		gin.SetMode(gin.DebugMode)
	}
	handler := gin.New()
	v1.NewRouter(handler, log, v1.UseCases{
		ExampleUseCase: exampleUseCase,
	})
	httpServer := httpserver.New(handler, ((*httpserver.Config)(&cfg.Http)))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Fatal(log, "app - Run - httpServer.Notify", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Fatal(log, "app - Run - httpServer.Shutdown", err)
	}
}
