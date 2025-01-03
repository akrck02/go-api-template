package middleware

import (
	"time"

	"github.com/akrck02/godot-api-template/models"
)

func Trazability(context *models.ApiContext) *models.Error {

	time := time.Now().UnixMilli()

	context.Trazability = models.Trazability{
		Endpoint:  context.Trazability.Endpoint,
		Timestamp: &time,
	}

	return nil
}
