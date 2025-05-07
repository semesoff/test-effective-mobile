package main

import (
	_ "service/docs"
	"service/pkg/server"
)

func main() {
	var serverProvider server.Server = server.NewServerManager()
	serverProvider.Start()
}
