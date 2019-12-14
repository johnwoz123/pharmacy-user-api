package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Healthy(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}
