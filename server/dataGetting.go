package server

import (
	"github.com/Rayato159/nevers-kube/entities"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) DataGetting(c echo.Context) error {
	resp := new(entities.Image)

	key := c.Param("key")

	result := s.db.First(resp).Where("id = ?", key)

	if result.Error != nil {
		s.logger.Errorf("Error getting image: %s", result.Error.Error())
		return result.Error
	}

	return c.JSON(200, resp)
}
