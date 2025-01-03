package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/akrck02/godot-api-template/models"
)

const CONTENT_TYPE_HEADER = "Content-Type"

// This struct represents an empty response
type EmptyResponse struct{}

// Response handle middleware function
func Response(context *models.ApiContext, writer http.ResponseWriter) {

	switch context.Trazability.Endpoint.ResponseMimeType {
	case models.MimeApplicationJson:
		sendJsonCatchingErrors(context, writer)
	default:
		sendResponseCatchingErrors(context, writer)
	}
}

// Send HTTP response
func SendResponse(w http.ResponseWriter, status int, response interface{}, contentType models.MimeType) {
	w.Header().Set(CONTENT_TYPE_HEADER, string(contentType))
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Send json data catching errors
func sendJsonCatchingErrors(context *models.ApiContext, writer http.ResponseWriter) {

	// calculate the time of the request
	start := time.Now()

	// execute the function
	response, responseError := context.Trazability.Endpoint.Listener(context)

	// calculate the time of the response
	end := time.Now()
	elapsed := end.Sub(start)

	// if something went wrong, return error
	if nil != responseError {
		SendResponse(writer, responseError.Status, responseError, models.MimeApplicationJson)
		return
	}

	// if response is nil, return {}
	if nil == response {
		response = &models.Response{
			Code:     http.StatusNoContent,
			Response: EmptyResponse{},
		}
	}

	// send response
	response.ResponseTime = elapsed.Nanoseconds()
	context.Response = *response
	SendResponse(writer, response.Code, response, models.MimeApplicationJson)
}

// Send HTTP response catching errors
func sendResponseCatchingErrors(context *models.ApiContext, writer http.ResponseWriter) {
	// execute the function
	result, responseError := context.Trazability.Endpoint.Listener(context)

	// if something went wrong, return error
	if nil != responseError {
		SendResponse(writer, responseError.Status, responseError, models.MimeApplicationJson)
		return
	}

	// if response is nil, return nothing
	if nil == result {
		SendResponse(writer, http.StatusNoContent, nil, context.Trazability.Endpoint.ResponseMimeType)
		return
	}

	// send response
	context.Response = *result
	SendResponse(writer, result.Code, result.Response, context.Trazability.Endpoint.ResponseMimeType)
}
