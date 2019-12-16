package app

import (
	"github.com/gin-gonic/gin"
	"github.com/johnwoz123/pharmacy-user-api/log"
)

var (
	router = gin.Default()
)

func BootstrapApp() {
	// Create server
	mapRoutes()
	log.Log.Info("starting the application......")
	router.Run(":8080")
}
