package health_check_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Healthy(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}
