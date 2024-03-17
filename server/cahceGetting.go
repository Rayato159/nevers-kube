package server

import (
	"context"
	"encoding/json"

	"github.com/Rayato159/nevers-kube/entities"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) CacheGetting(c echo.Context) error {
	ctx := context.Background()

	key := c.Param("key")

	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	resp := new(entities.Image)
	if err := json.Unmarshal([]byte(val), &resp); err != nil {
		s.logger.Error(err)
		return c.String(500, err.Error())
	}

	return c.JSON(200, resp)
}
