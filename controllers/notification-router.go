package controllers

import (
	"main/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type INotificationController interface {
	Create(route *gin.Engine)
	SetAsRead(route *gin.Engine)
	GetById(route *gin.Engine)
	DeleteByUserAndIds(route *gin.Engine)
}

type NotificationController struct {
	notificationService services.INotificationService
}

type NotificationControllerDependencies struct {
	dig.In
	NotificationService services.INotificationService `name:"NotificationService"`
}

func NotificationControllerInstance(deps NotificationControllerDependencies) *NotificationController {
	return &NotificationController{
		notificationService: deps.NotificationService,
	}
}

func (instance *NotificationController) Create(route *gin.Engine) {

	route.POST("notifications/add", instance.notificationService.AddNotification)
}

func (instance *NotificationController) SetAsRead(route *gin.Engine) {

	route.PUT("notifications/:id/read", instance.notificationService.SetNotificationRead)
}

func (instance *NotificationController) GetById(route *gin.Engine) {
	route.GET("notifications/user/:userId", instance.notificationService.GetNotifications)
}

func (instance *NotificationController) DeleteByUserAndIds(route *gin.Engine) {
	route.DELETE("notifications/user/:userId", instance.notificationService.DropNotifications)
}
