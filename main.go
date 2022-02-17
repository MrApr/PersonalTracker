package main

import (
	"flag"
	"github.com/MrApr/PersonalTracker/config"
	"github.com/MrApr/PersonalTracker/models"
	"github.com/MrApr/PersonalTracker/server"
	"github.com/MrApr/PersonalTracker/services"
)

var migrate *bool = flag.Bool("migrate", false, "Migrate database and creates schema")

func main() {
	config.CreateNewReader("env").Load(".env")
	models.ConnectToDb(100, 1000)
	flag.Parse()
	if *migrate {
		models.Migrate()
		return
	}
	sv := server.ConfigureServer("localhost", 8000)

	sv.Get("/collections", services.GetCollections)
	sv.Get("/show", services.RenderCollections)
	sv.Post("/collections/create", services.CreateNewCollection)
	sv.Post("/collections/edit", services.UpdateCollection)
	sv.Post("/collections/delete", services.DeleteCollection)

	sv.Get("/tasks", services.GetTasks)
	sv.Post("/tasks/create", services.CreateTask)
	sv.Post("/tasks/edit", services.EditTask)
	sv.Post("/tasks/delete", services.DeleteTask)

	//Todo Dockerize project & readme md
	//Todo Add tests for services & repos
	//Todo Test with locust

	sv.StartServer()
}
