package app

import (
	"github.com/johnwoz123/pharmacy-user-api/controllers/health_check_controller"
	"github.com/johnwoz123/pharmacy-user-api/controllers/user_controller"
)

func mapRoutes() {
	// HEALTH CHECK
	router.GET("/health-check", health_check_controller.Healthy)
	router.GET("/internal/users/search", user_controller.Search)
	// USERS
	router.POST("/users", user_controller.CreateUser)
	router.PUT("/users/:user_id", user_controller.UpdateUser)
	router.GET("/users/:user_id", user_controller.GetUsers)
	router.DELETE("/users/:user_id", user_controller.DeleteUser)

	// INTERNAL Routes

	//router.GET("/users/search", controllers.FindUser)
}
