package config

import (
	"main/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.Request.RequestURI == "/auth/token" {
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

		isValid := services.ValidateToken(token)

		if !isValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Not authorized")
			return
		} else {
			ctx.Next()
		}
	}

}
