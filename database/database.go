package database

import (
	"fmt"
	"sync"

	"github.com/Rayato159/nevers-kube/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InstanceGetting(conf *config.DatabaseConfig) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("user:pass@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.Host,
			conf.Port,
			conf.DBName,
		)

		if conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			panic(err)
		} else {
			db = conn
		}
	})

	return db
}
