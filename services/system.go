package services

import (
	"net/http"

	"github.com/akrck02/godot-api-template/models"
)

func Health(context *models.ApiContext) (*models.Response, *models.Error) {
	return &models.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func EmptyCheck(context *models.ApiContext) *models.Error {
	return nil
}
