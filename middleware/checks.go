package middleware

import (
	"github.com/akrck02/godot-api-template/models"
)

func Checks(context *models.ApiContext) *models.Error {

	checkError := context.Trazability.Endpoint.Checks(context)
	if checkError != nil {
		return checkError
	}

	return nil
}
