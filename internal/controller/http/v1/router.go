package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	weeklyevent "github.com/vicdevcode/schedule/internal/controller/http/v1/weekly_event"
	"github.com/vicdevcode/schedule/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, uc usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/api/v1")
	{
		weeklyevent.New(h, uc.WeeklyEventUseCase, l)
	}
}
