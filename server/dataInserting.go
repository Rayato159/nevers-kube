package server

import (
	"github.com/Rayato159/nevers-kube/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) DataInserting(c echo.Context) error {
	reqImage := new(entities.Image)
	if err := c.Bind(reqImage); err != nil {
		return err
	}

	uuidV7, _ := uuid.NewV7()
	reqImage.ID = uuidV7.String()

	if err := s.db.Create(reqImage).Error; err != nil {
		return err
	}

	return c.JSON(200, reqImage)
}
