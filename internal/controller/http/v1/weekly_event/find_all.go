package weeklyevent

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicdevcode/schedule/internal/controller/http/v1/exception"
	"github.com/vicdevcode/schedule/internal/entity"
)

type findAllResponse struct {
	WeeklyEvents []entity.WeeklyEvent `json:"weekly_events"`
}

func (r *weeklyEventRoutes) findAll(c *gin.Context) {
	weeklyEvents, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error("http-v1-weekly-events-find-all", slog.Any("error", err))
		exception.Response(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, findAllResponse{weeklyEvents})
}
