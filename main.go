package main

import "github.com/MrApr/PersonalTracker/server"

func main() {
	server.ConfigureServer("localhost", 8000).StartServer()
}
