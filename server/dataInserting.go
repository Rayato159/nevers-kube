package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/Rayato159/nevers-kube/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) DataInserting(c echo.Context) error {
	ctx := context.Background()
	key := c.Param("key")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		s.logger.Error("get file error: ", err.Error())
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		s.logger.Error("get file error: ", err.Error())
		return err
	}

	defer file.Close()

	var data []byte
	_, err = file.Read(data)
	if err != nil {
		s.logger.Error("get file error: ", err.Error())
		return err
	}

	reqImage := new(entities.Image)
	uuidV7, _ := uuid.NewV7()
	reqImage.ID = uuidV7.String()
	reqImage.ImageBase64 = base64.StdEncoding.EncodeToString(data)

	reqImageJson, err := json.Marshal(reqImage)
	if err != nil {
		s.logger.Error("Error sequenced json data: ", err.Error())
		return err
	}

	if result := s.rdb.Set(ctx, key, string(reqImageJson), 5*time.Minute); result.Err() != nil {
		s.logger.Errorf("Error setting cache: %s", result.Err().Error())
		return result.Err()
	}

	if err := s.db.Create(reqImage).Error; err != nil {
		s.logger.Errorf("Error creating image: %s", err.Error())
		return err
	}

	return c.JSON(200, reqImage)
}
