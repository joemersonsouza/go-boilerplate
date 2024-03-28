package controllers

import (
	"main/services"

	"github.com/gin-gonic/gin"
)

func NotificationController(route *gin.Engine) {

	route.POST("notifications/add", services.AddNotification)

	route.PUT("notifications/:id/read", services.SetNotificationRead)

	route.GET("notifications/user/:userId", services.GetNotifications)

	route.DELETE("notifications/user/:userId", services.DropNotifications)
}
