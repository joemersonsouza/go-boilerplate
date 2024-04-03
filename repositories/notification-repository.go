package repository

import (
	"errors"

	"github.com/google/uuid"
	"go.uber.org/dig"
)

type INotificationRepository interface {
	Insert(notification NotificationObject) error
	GetByUserIdAndState(userId string, state string) ([]NotificationObject, error)
	SetStateRead(notificationId string) error
	DeleteIdIn(ids string) error
}

type NotificationRepository struct {
	database IDatabase
}

type NotificationRepositoryDependencies struct {
	dig.In
	Database IDatabase `name:"Database"`
}

type NotificationObject struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	CompanyId string `json:"companyId"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
}

func NotificationRepositoryInstance(deps NotificationRepositoryDependencies) *NotificationRepository {
	return &NotificationRepository{
		database: deps.Database,
	}
}

func (instance *NotificationRepository) Insert(notification NotificationObject) error {
	Db := instance.database.ConnectDatabase()

	id := uuid.New().String()

	_, err := Db.Exec("insert into notification(id, user_id,company_id,message) values ($1,$2,$3,$4)", id, notification.UserId, notification.CompanyId, notification.Message)

	return err
}

func (instance *NotificationRepository) GetByUserIdAndState(userId string, state string) ([]NotificationObject, error) {
	Db := instance.database.ConnectDatabase()

	query := "select id, user_id, company_id, message, read from notification where user_id = $1"

	if state == "READ" {
		query = query + " and read = true"
	} else if state == "UNREAD" {
		query = query + " and read = false"
	}

	rows, err := Db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	var notifications []NotificationObject

	for rows.Next() {
		var notification NotificationObject
		rows.Scan(&notification.Id, &notification.UserId, &notification.CompanyId, &notification.Message, &notification.Read)
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (instance *NotificationRepository) SetStateRead(notificationId string) error {

	Db := instance.database.ConnectDatabase()

	rows, err := Db.Query("select 1 from notification where id = $1", notificationId)

	if err != nil {
		return err
	}

	var hasData int = 0

	for rows.Next() {
		rows.Scan(&hasData)
	}

	if hasData == 0 {
		return errors.New("no data")
	}

	_, err = Db.Exec("update notification set read = true where id = $1", notificationId)

	if err != nil {
		return err
	}

	return nil
}

func (instance *NotificationRepository) DeleteIdIn(ids string) error {
	Db := instance.database.ConnectDatabase()

	_, err := Db.Exec("delete from notification where id in ($1)", ids)

	return err
}
