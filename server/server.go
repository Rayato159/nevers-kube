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
	"gorm.io/gorm"
)

type echoServer struct {
	app    *echo.Echo
	conf   *config.ServerConfig
	db     *gorm.DB
	logger echo.Logger
}

func ServerInstaceGetting(conf *config.ServerConfig, db *gorm.DB) *echoServer {
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)

	logger := app.Logger
	return &echoServer{app, conf, db, logger}
}

func (s *echoServer) Starting() {
	router := s.app.Group("/v1")

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	router.GET("/cache", s.CacheGetting)
	router.DELETE("/cache", s.CacheClearing)

	router.POST("/data", s.DataInserting)
	router.GET("/data", s.DataGetting)

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
