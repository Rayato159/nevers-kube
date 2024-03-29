package database

import (
	"fmt"
	"sync"

	"github.com/Rayato159/nevers-kube/config"
	"github.com/Rayato159/nevers-kube/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InstanceGetting(conf *config.DatabaseConfig) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
		)

		if conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			panic(err)
		} else {
			db = conn
		}

		db.AutoMigrate(&entities.Image{})
	})

	return db
}
