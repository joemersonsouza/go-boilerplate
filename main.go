package main

import (
	"log"
	"main/config"
	"main/server"
)

func main() {
	container := config.Inject()
	err := container.Invoke(server.StartServer)

	if err != nil {
		log.Fatal(err)
	}
}
