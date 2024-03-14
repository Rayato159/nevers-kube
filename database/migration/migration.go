package migration

import (
	"github.com/Rayato159/nevers-kube/config"
	"github.com/Rayato159/nevers-kube/database"
	"github.com/Rayato159/nevers-kube/entities"
)

func main() {
	conf := config.InstaceGetting()

	db := database.InstanceGetting(conf.DatabaseConfig)

	db.Migrator().CreateTable(&entities.Image{})
}
