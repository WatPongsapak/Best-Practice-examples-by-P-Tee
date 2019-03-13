package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Endpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
