package config

import (
	"main/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type IAuthenticator interface {
	TokenAuthMiddleware() gin.HandlerFunc
}

type Authenticator struct {
	authenticationService services.IAuthenticationService
}

type AuthenticatorDependencies struct {
	dig.In
	AuthenticationService services.IAuthenticationService `name:"AuthenticationService"`
}

func AuthenticatorInstance(deps AuthenticatorDependencies) *Authenticator {
	return &Authenticator{
		authenticationService: deps.AuthenticationService,
	}
}

func (instance *Authenticator) TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.Request.RequestURI == "/auth/token" || ctx.Request.RequestURI == "/health/status" {
			ctx.Next()
			return
		}

		tokenHeader := ctx.Request.Header["Authorization"]

		if len(tokenHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Not authorized")
			return
		}

		token := tokenHeader[0]
		token = strings.Replace(token, "Bearer ", "", -1)

		isValid := instance.authenticationService.ValidateToken(token)

		if !isValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Not authorized")
			return
		} else {
			ctx.Next()
		}
	}

}
