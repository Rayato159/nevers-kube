package server

import (
	"github.com/Rayato159/nevers-kube/entities"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) DataGetting(c echo.Context) error {
	var resp *entities.Image

	key := c.QueryParam("key")

	result := s.db.First(resp).Where("id = ?", key)

	if result.Error != nil {
		return result.Error
	}

	return c.JSON(200, resp)
}
