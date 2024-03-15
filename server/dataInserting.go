package server

import (
	"context"
	"time"

	"github.com/Rayato159/nevers-kube/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) DataInserting(c echo.Context) error {
	ctx := context.Background()

	reqImage := new(entities.Image)
	if err := c.Bind(reqImage); err != nil {
		return err
	}

	uuidV7, _ := uuid.NewV7()
	reqImage.ID = uuidV7.String()

	if result := s.rdb.Set(ctx, reqImage.ID, reqImage.ImageBase64, 5*time.Minute); result.Err() != nil {
		s.logger.Errorf("Error setting cache: %s", result.Err().Error())
		return result.Err()
	}

	if err := s.db.Create(reqImage).Error; err != nil {
		s.logger.Errorf("Error creating image: %s", err.Error())
		return err
	}

	return c.JSON(200, reqImage)
}
