package controllers

import (
	"main/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type IAuthenticationController interface {
	CreateToken(route *gin.Engine)
}

type AuthenticationController struct {
	authenticationService services.IAuthenticationService
}

type AuthenticationControllerDependencies struct {
	dig.In
	AuthenticationService services.IAuthenticationService `name:"AuthenticationService"`
}

func AuthenticationControllerInstance(deps AuthenticationControllerDependencies) *AuthenticationController {
	return &AuthenticationController{
		authenticationService: deps.AuthenticationService,
	}
}

func (instance *AuthenticationController) CreateToken(route *gin.Engine) {
	route.POST("auth/token", instance.authenticationService.CreateToken)
}
