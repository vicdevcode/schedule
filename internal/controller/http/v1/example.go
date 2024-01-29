package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/init_template/internal/entity"
	"github.com/vicdevcode/init_template/internal/usecase"
)

type exampleRoutes struct {
	u usecase.Example
	l *slog.Logger
}

func newExampleRoutes(handler *gin.RouterGroup, u usecase.Example, l *slog.Logger) {
	r := &exampleRoutes{u, l}
	h := handler.Group("/example")
	{
		h.GET("/", r.getExamples)
		h.POST("/", r.createExample)
	}
}

type getExamplesResponse struct {
	Examples []entity.Example `json:"examples"`
}

func (r *exampleRoutes) getExamples(c *gin.Context) {
	examples, err := r.u.GetExamples(c.Request.Context())
	if err != nil {
		r.l.Error("http-v1-get-example", slog.Any("error", err))
		errorResponse(c, http.StatusInternalServerError, "can not get examples")
		return
	}
	c.JSON(http.StatusOK, getExamplesResponse{examples})
}

type createExampleResponse struct {
	Example *entity.Example `json:"example"`
}

type createExampleRequest struct {
	Message string `json:"message"`
}

func (r *exampleRoutes) createExample(c *gin.Context) {
	var body createExampleRequest

	if err := c.BindJSON(&body); err != nil {
		return
	}

	example, err := r.u.CreateExample(c.Request.Context(), body.Message)
	if err != nil {
		r.l.Error("http-v1-get-example", slog.Any("error", err))
		errorResponse(c, http.StatusInternalServerError, "can not create example")
		return
	}
	c.JSON(http.StatusOK, createExampleResponse{example})
}
