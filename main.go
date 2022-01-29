package main

import (
	"flag"
	"github.com/MrApr/PersonalTracker/repositories"
	"github.com/MrApr/PersonalTracker/server"
)

var migrate *bool = flag.Bool("migrate", false, "Migrate database and creates schema")

func main() {
	flag.Parse()
	if *migrate {
		repositories.Migrate()
		return
	}
	server.ConfigureServer("localhost", 8000).StartServer()
}
