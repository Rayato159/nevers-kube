package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rayato159/nevers-kube/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type echoServer struct {
	app    *echo.Echo
	conf   *config.ServerConfig
	db     *gorm.DB
	rdb    *redis.Client
	logger echo.Logger
}

func ServerInstaceGetting(conf *config.ServerConfig, db *gorm.DB, rdb *redis.Client) *echoServer {
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)

	logger := app.Logger
	return &echoServer{app, conf, db, rdb, logger}
}

func (s *echoServer) Starting() {
	router := s.app.Group("/api/v1")

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	s.app.Use(middleware.BodyLimit(s.conf.BodyLimit))

	s.app.GET("/health", func(c echo.Context) error {
		s.logger.Info("Health check request received")
		err := c.String(http.StatusOK, "OK")
		if err != nil {
			s.logger.Errorf("Error responding to health check: %s", err.Error())
		} else {
			s.logger.Error("Health check response sent successfully")
		}
		return err
	})

	router.GET("/cache/:key", s.CacheGetting)
	router.DELETE("/cache", s.CacheClearing)

	router.POST("/data", s.DataInserting)
	router.GET("/data/:key", s.DataGetting)

	router.POST("/machine", s.MachineStressTesting)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefulShutdown(quit)

	s.httpListening()
}

func (s *echoServer) httpListening() {
	uri := fmt.Sprintf(":%d", s.conf.Port)

	if err := s.app.Start(uri); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %v", err)
	}
}

func (s *echoServer) gracefulShutdown(quit <-chan os.Signal) {
	ctx := context.Background()

	<-quit
	s.app.Logger.Infof("Shutting down service...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}
