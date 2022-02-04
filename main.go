package main

import (
	"flag"
	"github.com/MrApr/PersonalTracker/models"
	"github.com/MrApr/PersonalTracker/server"
)

var migrate *bool = flag.Bool("migrate", false, "Migrate database and creates schema")

func main() {
	models.ConnectToDb(100, 1000)
	flag.Parse()
	if *migrate {
		models.Migrate()
		return
	}
	server.ConfigureServer("localhost", 8000).StartServer()
}
