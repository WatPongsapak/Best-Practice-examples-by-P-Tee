package merchant

import (
	"context"
	"net/http"

	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/config"

	"github.com/gin-gonic/gin"
)

type merchantService interface {
	Register(ctx context.Context, name string) (RegisterResponse, error)
	Information(ctx context.Context, id string) (Merchant, error)
}

type Handler struct {
	conf        *config.Config
	merchantSrv merchantService
}

func NewHandler(conf *config.Config, srv merchantService) *Handler {
	return &Handler{conf: conf, merchantSrv: srv}
}

func (h Handler) Register(c *gin.Context) {
	if !h.conf.Merchant.Enable {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.merchantSrv.Register(c, req.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h Handler) Information(c *gin.Context) {
	if !h.conf.Merchant.Enable {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var req struct {
		ID string `json:"id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.merchantSrv.Information(c, req.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
