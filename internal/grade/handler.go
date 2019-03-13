package grade

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type gradeService interface {
	Calculate(ctx context.Context, score string) (CalculateResponse, error)
}

type Handler struct {
	gradeSrv gradeService
}

func NewHandler(srv gradeService) *Handler {
	return &Handler{gradeSrv: srv}
}

func (h Handler) Calculate(c *gin.Context) {
	s := c.Request.URL.Query().Get("score")
	res, err := h.gradeSrv.Calculate(c, s)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
