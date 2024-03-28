package main

import (
	"embed"
	"main/config"
	"main/controllers"

	"github.com/gin-gonic/gin"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	// gin.SetMode(gin.ReleaseMode) //optional to not get warning
	// route.SetTrustedProxies([]string{"192.168.1.2"}) //to trust only a specific value
	route := gin.Default()
	route.Use(config.TokenAuthMiddleware())

	config.RunMigration(embedMigrations)

	controllers.NotificationController(route)
	controllers.HealthController(route)
	controllers.AuthenticationController(route)

	err := route.Run(":8080")

	if err != nil {
		panic(err)
	}

}
