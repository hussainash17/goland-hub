package routes

import "github.com/gin-gonic/gin"

import (
	"awesomeProject/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API routes
	router.GET("/health", controllers.HealthCheck)
	router.GET("/users", controllers.GetUsers) // Example endpoint
	router.POST("/users", controllers.CreateUser)

	return router
}
