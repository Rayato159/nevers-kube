package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var once sync.Once

type (
	DatabaseConfig struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
	}

	ServerConfig struct {
		Port         int      `mapstructure:"port" validate:"required"`
		AllowOrigins []string `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string   `mapstructure:"bodyLimit" validate:"required"`
	}

	AppConfig struct {
		DatabaseConfig *DatabaseConfig `mapstructure:"database" validate:"required"`
		ServerConfig   *ServerConfig   `mapstructure:"server" validate:"required"`
	}
)

var appConfigInstance *AppConfig

func InstaceGetting() *AppConfig {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./etc")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("read config file failed: %v", err))
		}

		if err := viper.Unmarshal(&appConfigInstance); err != nil {
			panic(fmt.Errorf("unmarshalkey config file failed: %v", err))
		}

		validate := validator.New()

		if err := validate.Struct(appConfigInstance); err != nil {
			panic(fmt.Errorf("validate config file failed: %v", err))
		}
	})

	return appConfigInstance
}
