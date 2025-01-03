package middleware

import (
	"net/http"

	"github.com/akrck02/godot-api-template/models"
)

const USER_AGENT_HEADER = "User-Agent"

func Request(r *http.Request, context *models.ApiContext) *models.Error {

	parserError := r.ParseForm()

	if parserError != nil {
		return &models.Error{
			Status:  http.StatusBadRequest,
			Error:   0,
			Message: parserError.Error(),
		}
	}

	context.Request = models.Request{
		Authorization: r.Header.Get(AUTHORITATION_HEADER),
		Ip:            r.Host,
		UserAgent:     r.Header.Get(USER_AGENT_HEADER),
		//Headers:       r.Header,
		Body: r.Body,
		//Params:        r.Form,
	}

	// If files are present, add them to the context
	if r.MultipartForm != nil {
		context.Request.Files = r.MultipartForm.File
	}

	return nil
}
