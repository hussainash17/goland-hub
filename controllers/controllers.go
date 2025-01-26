package controllers

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}

func GetUsers(context *gin.Context) {
	var news []models.News
	result := config.DB.Find(&news)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the news slice directly instead of the result object
	context.JSON(http.StatusOK, gin.H{
		"data":  news,
		"count": len(news),
	})
}

func CreateUser(context *gin.Context) {

}
