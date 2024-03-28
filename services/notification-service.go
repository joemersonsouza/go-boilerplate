package services

import (
	"encoding/json"
	"fmt"
	repository "main/repositories"
	"main/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jrivets/log4g"
)

type NotificationObject struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	CompanyId string `json:"companyId"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
}

func AddNotification(ctx *gin.Context) {
	var body NotificationObject
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

	repository.ConnectDatabase()
	id := uuid.New().String()

	_, err = repository.Db.Exec("insert into notification(id, user_id,company_id,message) values ($1,$2,$3,$4)", id, body.UserId, body.CompanyId, body.Message)

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(400, "Something went wrong, review your request and try again")
	} else {
		ctx.JSON(http.StatusOK, "Notification is successfully created.")
		logger.Info(fmt.Sprintf("The notification %s was created", id))
	}

}

func GetNotifications(ctx *gin.Context) {
	logger := log4g.GetLogger(util.LoggerName)
	userId := ctx.Param("userId")
	state := ctx.Query("state")
	query := "select id, user_id, company_id, message, read from notification where user_id = $1"

	repository.ConnectDatabase()

	if state == "READ" {
		query = query + " and read = true"
	} else if state == "UNREAD" {
		query = query + " and read = false"
	}

	rows, err := repository.Db.Query(query, userId)

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	var notifications []NotificationObject

	for rows.Next() {
		var notification NotificationObject
		rows.Scan(&notification.Id, &notification.UserId, &notification.CompanyId, &notification.Message, &notification.Read)
		notifications = append(notifications, notification)
	}

	if notifications == nil {
		ctx.AbortWithStatusJSON(404, "Not found")
		return
	}

	ctx.IndentedJSON(http.StatusOK, notifications)
}

func SetNotificationRead(ctx *gin.Context) {
	logger := log4g.GetLogger(util.LoggerName)
	notificationId := ctx.Param("id")
	repository.ConnectDatabase()

	rows, err := repository.Db.Query("select 1 from notification where id = $1", notificationId)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	var hasData int = 0

	for rows.Next() {
		rows.Scan(&hasData)
	}

	if hasData == 0 {
		ctx.AbortWithStatusJSON(404, "Notication does not exist")
		return
	}

	_, err = repository.Db.Exec("update notification set read = true where id = $1", notificationId)

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(500, "Something went wrong")
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"result": "Done"})
}

func DropNotifications(ctx *gin.Context) {
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

	repository.ConnectDatabase()

	_, err = repository.Db.Exec("delete from notification where id in ($1)", strings.Join(body, ","))

	if err != nil {
		logger.Error(err.Error())
		ctx.AbortWithStatusJSON(404, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"result": "Done"})
}
