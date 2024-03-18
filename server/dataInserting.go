package server

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Rayato159/nevers-kube/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) DataInserting(c echo.Context) error {
	ctx := context.Background()

	reqImage := new(entities.Image)
	if err := c.Bind(reqImage); err != nil {
		s.logger.Error("Error binding request: ", err.Error())
		return c.String(500, err.Error())
	}

	uuidV7, _ := uuid.NewV7()
	reqImage.ID = uuidV7.String()

	reqImageJson, err := json.Marshal(reqImage)
	if err != nil {
		s.logger.Error("Error sequenced json data: ", err.Error())
		return err
	}

	if result := s.rdb.Set(ctx, reqImage.ID, string(reqImageJson), 5*time.Minute); result.Err() != nil {
		s.logger.Errorf("Error setting cache: %s", result.Err().Error())
		return result.Err()
	}

	if err := s.db.Create(reqImage).Error; err != nil {
		s.logger.Errorf("Error creating image: %s", err.Error())
		return err
	}

	return c.JSON(201, reqImage)
}
