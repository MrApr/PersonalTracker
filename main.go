package main

import (
	"flag"
	"github.com/MrApr/PersonalTracker/Repositories"
	"github.com/MrApr/PersonalTracker/server"
)

var migrate *bool = flag.Bool("migrate", false, "Migrate database and creates schema")

func main() {
	flag.Parse()
	if *migrate {
		Repositories.Migrate()
		return
	}
	server.ConfigureServer("localhost", 8000).StartServer()
}
