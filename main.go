package main

import (
	"flag"
	"github.com/MrApr/PersonalTracker/models"
	"github.com/MrApr/PersonalTracker/server"
	"github.com/MrApr/PersonalTracker/services"
)

var migrate *bool = flag.Bool("migrate", false, "Migrate database and creates schema")

func main() {
	models.ConnectToDb(100, 1000)
	flag.Parse()
	if *migrate {
		models.Migrate()
		return
	}
	sv := server.ConfigureServer("localhost", 8000)

	sv.Get("/collections", services.GetCollections)
	sv.Post("/collections/create", services.CreateNewCollection)
	sv.Post("/collections/edit", services.UpdateCollection)
	sv.Post("/collections/delete", services.DeleteCollection)

	sv.Get("/tasks", services.GetTasks)
	sv.Post("/tasks/create", services.CreateTask)
	sv.Post("/tasks/edit", services.EditTask)
	sv.Post("/tasks/delete", services.DeleteTask)

	//Todo add Template layer
	//Todo Add logging layer (both file and remote)
	//Todo make it to read from conf file
	//Todo define custom error types for better err handling
	//Todo Dockerize project
	//Todo Add tests for services & repos
	//Todo Test with locust

	sv.StartServer()
}
