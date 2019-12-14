package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!!!")
}
func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!!!")
}
//func FindUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!!!")
//}
