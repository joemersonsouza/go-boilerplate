package controllers

import (
	"main/services"

	"github.com/gin-gonic/gin"
)

func AuthenticationController(route *gin.Engine) {

	route.POST("auth/token", services.CreateToken)
}
