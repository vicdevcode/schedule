package weeklyevent

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicdevcode/schedule/internal/controller/http/v1/exception"
	"github.com/vicdevcode/schedule/internal/entity"
	"github.com/vicdevcode/schedule/internal/model"
)

type createRequest struct {
	model.CreateWeeklyEvent
}

type createResponse struct {
	WeeklyEvent *entity.WeeklyEvent `json:"weekly_event"`
}

func (r *weeklyEventRoutes) create(c *gin.Context) {
	var body createRequest

	if err := c.BindJSON(&body); err != nil {
		return
	}

	weeklyEvent, err := r.u.Create(c.Request.Context(), body.CreateWeeklyEvent)
	if err != nil {
		r.l.Error("http-v1-weekly-event-create", slog.Any("error", err))
		exception.Response(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createResponse{WeeklyEvent: weeklyEvent})
}
