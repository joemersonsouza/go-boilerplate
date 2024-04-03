package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthController interface {
	Status(route *gin.Engine)
}

type HealthController struct{}

func HealthControllerInstance() *HealthController {
	return &HealthController{}
}

func (instance *HealthController) Status(route *gin.Engine) {

	route.GET("/health/status", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}
