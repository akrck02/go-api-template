package middleware

import (
	"github.com/akrck02/godot-api-template/models"
)

type Middleware func(context *models.ApiContext) *models.Error
