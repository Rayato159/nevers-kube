package main

import (
	"github.com/Rayato159/nevers-kube/config"
	"github.com/Rayato159/nevers-kube/server"
)

func main() {
	conf := config.InstaceGetting()

	// db := database.InstanceGetting(conf.DatabaseConfig)

	s := server.ServerInstaceGetting(conf.ServerConfig, nil)
	s.Starting()
}
