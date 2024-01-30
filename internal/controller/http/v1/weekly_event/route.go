package weeklyevent

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/init_template/internal/usecase"
)

type response struct {
	Error string `json:"error"`
}

type weeklyEventRoutes struct {
	u usecase.WeeklyEvent
	l *slog.Logger
}

func New(handler *gin.RouterGroup, u usecase.WeeklyEvent, l *slog.Logger) {
	r := &weeklyEventRoutes{u, l}
	h := handler.Group("/weekly-event")
	{
		h.GET("/", r.findAll)
		h.POST("/", r.create)
		h.DELETE("/", r.deleteEntity)
	}
}
