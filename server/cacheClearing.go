package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *echoServer) CacheClearing(c echo.Context) error {
	ctx := context.Background()

	if err := s.rdb.FlushAll(ctx).Err(); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Cache Cleared!")
}
