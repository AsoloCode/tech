package main

import (
	"tech/cmd/server"
	"tech/pkg/config"
)

func main() {
	// setup various configuration for app
	config.LoadAllConfigs(".env")

	server.Server()
}
