package services

import (
	"net/http"

	"github.com/akrck02/godot-api-template/errors"
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

func NotImplemented(context *models.ApiContext) (*models.Response, *models.Error) {

	return nil, &models.Error{
		Error:   errors.NotImplemented,
		Message: "Not implemented",
		Status:  http.StatusNotImplemented,
	}
}
