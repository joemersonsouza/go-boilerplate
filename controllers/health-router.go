package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthController(route *gin.Engine) {

	route.GET("/health/status", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}
