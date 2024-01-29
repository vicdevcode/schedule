package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/init_template/internal/usecase"
)

type UseCases struct {
	ExampleUseCase usecase.Example
}

func NewRouter(handler *gin.Engine, l *slog.Logger, uc UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/api/v1")
	{
		newExampleRoutes(h, uc.ExampleUseCase, l)
	}
}
