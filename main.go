package main

import (
	"github.com/Rayato159/nevers-kube/cache"
	"github.com/Rayato159/nevers-kube/config"
	"github.com/Rayato159/nevers-kube/database"
	"github.com/Rayato159/nevers-kube/server"
	_ "go.uber.org/automaxprocs"
)

func main() {
	conf := config.InstaceGetting()

	db := database.InstanceGetting(conf.DatabaseConfig)

	rdb := cache.ExampleClient()

	s := server.ServerInstaceGetting(conf.ServerConfig, db, rdb)
	s.Starting()
}
