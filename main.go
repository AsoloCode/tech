package main

import (
	"tech/cmd/server"
	"tech/pkg/config"
)

func main() {

	config.LoadAllConfigs(".env")

	server.Server()
}
