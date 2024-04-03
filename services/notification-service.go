package services

import (
	"encoding/json"
	"fmt"
	repository "main/repositories"
	"main/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jrivets/log4g"
	"go.uber.org/dig"
)

type INotificationService interface {
	AddNotification(ctx *gin.Context)
	GetNotifications(ctx *gin.Context)
	SetNotificationRead(ctx *gin.Context)
	DropNotifications(ctx *gin.Context)
}

type NotificationService struct {
	notificationRepository repository.INotificationRepository
}

type NotificationServiceDependencies struct {
	dig.In
	NotificationRepository repository.INotificationRepository `name:"NotificationRepository"`
}

func NotificationServiceInstance(deps NotificationServiceDependencies) *NotificationService {
	return &NotificationService{
		notificationRepository: deps.NotificationRepository,
	}
}

func (instance *NotificationService) AddNotification(ctx *gin.Context) {
	var body repository.NotificationObject
	data, err := ctx.GetRawData()
	logger := log4g.GetLogger(util.LoggerName)

	if err != nil {
		ctx.AbortWithStatusJSON(400, "Notification is not defined")
		return
	}

	err = json.Unmarshal(data, &body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	err = instance.notificationRepository.Insert(body)

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Something went wrong, review your request and try again")
	} else {
		ctx.JSON(http.StatusOK, "Notification is successfully created.")
		logger.Info(fmt.Sprintf("The notification %s was created", body.Message))
	}

}

func (instance *NotificationService) GetNotifications(ctx *gin.Context) {
	logger := log4g.GetLogger(util.LoggerName)
	userId := ctx.Param("userId")
	state := ctx.Query("state")

	notifications, err := instance.notificationRepository.GetByUserIdAndState(userId, state)

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	if notifications == nil {
		ctx.AbortWithStatusJSON(404, "Not found")
		return
	}

	ctx.IndentedJSON(http.StatusOK, notifications)
}

func (instance *NotificationService) SetNotificationRead(ctx *gin.Context) {
	notificationId := ctx.Param("id")

	err := instance.notificationRepository.SetStateRead(notificationId)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"result": "Done"})
}

func (instance *NotificationService) DropNotifications(ctx *gin.Context) {
	logger := log4g.GetLogger(util.LoggerName)
	var body []string
	data, err := ctx.GetRawData()

	if err != nil {
		ctx.AbortWithStatusJSON(400, "Notification is not defined")
		return
	}

	err = json.Unmarshal(data, &body)

	if err != nil || len(body) == 0 {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	err = instance.notificationRepository.DeleteIdIn(strings.Join(body, ","))

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"result": "Done"})
}
