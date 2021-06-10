package main

import (
	config "Dig/ConfigModel/Config"
	"fmt"
	"log"

	"go.uber.org/dig"
)

func Builder() *dig.Container {
	container := dig.New()
	container.Provide(config.NewDataBaseConfig)
	container.Provide(config.ConnectDatabase)
	container.Provide(config.NewPersonRepository)
	container.Provide(config.NewServer)

	return container
}

func main() {
	fmt.Println("===== Dig POC =====")
	container := Builder()
	err := container.Invoke(func(server *config.Server) {
		server.Run()
	})
	if err != nil {
		log.Fatal("Error occur while start or Run Server..")
	}
}
