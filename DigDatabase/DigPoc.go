package main

import (
	config "DigDatabase/DatabaseConfig"
	repo "DigDatabase/Repository"
	controller "DigDatabase/Resources"
	core "DigDatabase/Service"
	"fmt"
	"log"

	"go.uber.org/dig"
)

func ConnectionBuilder() *dig.Container {
	container := dig.New()
	container.Provide(config.NewDatabaseConnection)
	container.Provide(config.NewRepositoryConnection)
	container.Provide(repo.NewUserRepo)
	container.Provide(repo.NewExpertRepo)
	container.Provide(core.NewService)
	container.Provide(controller.NewServer)
	return container
}

func main() {
	fmt.Println("===-== POC Regarding to DIG With Database Collection =====")
	log.Println("It's Time to Build a Container...")
	container := ConnectionBuilder()
	err := container.Invoke(func(service *controller.Server) {
		service.Run()
	})
	if err != nil {
		log.Fatal("Error Occur ::: ", err)
	}

}
