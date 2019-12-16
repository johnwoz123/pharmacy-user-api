package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func BootstrapApp() {
	// Create server
	mapRoutes()
	router.Run(":8080")
}
