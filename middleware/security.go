package middleware

import (
	"log"
	"net/http"

	"github.com/akrck02/godot-api-template/errors"
	"github.com/akrck02/godot-api-template/models"
)

const AUTHORITATION_HEADER = "Authorization"

func Security(context *models.ApiContext) *models.Error {

	// Check if endpoint is registered and secured
	if !context.Trazability.Endpoint.Secured {
		log.Printf("Endpoint %s is not secured", context.Trazability.Endpoint.Path)
		return nil
	}

	log.Printf("Endpoint %s is secured", context.Trazability.Endpoint.Path)

	// Check if token is empty
	if context.Request.Authorization == "" {
		return &models.Error{
			Status:  http.StatusForbidden,
			Error:   errors.InvalidToken,
			Message: "Missing token",
		}
	}

	// Check if token is valid
	if !tokenIsValid(context.Request.Authorization) {
		return &models.Error{
			Status:  http.StatusForbidden,
			Error:   errors.InvalidToken,
			Message: "Invalid token",
		}
	}

	return nil
}

// Check if token is valid
func tokenIsValid(_ string) bool {
	return true
}
