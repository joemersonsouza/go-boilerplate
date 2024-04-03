package server

import (
	"embed"
	"main/config"
	"main/controllers"
	repository "main/repositories"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Server struct {
	authenticator            config.IAuthenticator
	database                 repository.IDatabase
	authenticationController controllers.IAuthenticationController
	healthController         controllers.IHealthController
	notificationController   controllers.INotificationController
}

type ServerDependencies struct {
	dig.In

	Authenticator            config.IAuthenticator                 `name:"Authenticator"`
	Database                 repository.IDatabase                  `name:"Database"`
	AuthenticationController controllers.IAuthenticationController `name:"AuthenticationController"`
	HealthController         controllers.IHealthController         `name:"HealthController"`
	NotificationController   controllers.INotificationController   `name:"NotificationController"`
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func StartServer(deps ServerDependencies) {
	// gin.SetMode(gin.ReleaseMode) //optional to not get warning
	// route.SetTrustedProxies([]string{"192.168.1.2"}) //to trust only a specific value

	app := Server{
		authenticator:            deps.Authenticator,
		database:                 deps.Database,
		authenticationController: deps.AuthenticationController,
		healthController:         deps.HealthController,
		notificationController:   deps.NotificationController,
	}

	route := gin.Default()
	route.Use(app.authenticator.TokenAuthMiddleware())

	Db := app.database.ConnectDatabase()

	// Run migration
	app.database.RunMigration(embedMigrations, Db)

	//Expose controllers
	app.notificationController.Create(route)
	app.notificationController.DeleteByUserAndIds(route)
	app.notificationController.GetById(route)
	app.notificationController.SetAsRead(route)
	app.healthController.Status(route)
	app.authenticationController.CreateToken(route)

	err := route.Run(":8080")

	if err != nil {
		panic(err)
	}
}
