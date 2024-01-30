package weeklyevent

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicdevcode/schedule/internal/controller/http/v1/exception"
)

type deleteResponse struct {
	Message string `json:"message"`
}

func (r *weeklyEventRoutes) deleteEntity(c *gin.Context) {
	if err := r.u.Delete(c.Request.Context(), c.Query("id")); err != nil {
		r.l.Error("http-v1-weekly-event-remove", slog.Any("error", err))
		exception.Response(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, deleteResponse{Message: "weekly_event was deleted"})
}
