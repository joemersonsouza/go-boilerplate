package services

import (
	"fmt"
	repository "main/repositories"
	"main/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jrivets/log4g"
)

var secretKey = []byte("secret-key")

func CreateToken(ctx *gin.Context) {
	logger := log4g.GetLogger(util.LoggerName)
	apiKey := ctx.Request.Header["Api-Key"][0]

	company := repository.GetCompany(apiKey)

	if len(company.ApiKey) == 0 {
		logger.Error(fmt.Sprintf("Login attempt with wrong api key %s", apiKey))
		ctx.AbortWithStatusJSON(401, "Not authorized")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"companyId": company.Id,
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func ValidateToken(tokenString string) bool {
	logger := log4g.GetLogger(util.LoggerName)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Error validating token %s", err.Error()))
		return false
	}

	if !token.Valid {
		logger.Error(fmt.Sprintf("Usage attempt with wrong token %s", tokenString))
		return false
	}

	return true
}
