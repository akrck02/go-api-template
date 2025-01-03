package apicommon

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/akrck02/godot-api-template/configuration"
	"github.com/akrck02/godot-api-template/middleware"
	"github.com/akrck02/godot-api-template/models"
)

const API_PATH = "/api/"
const CONTENT_TYPE_HEADER = "Content-Type"

// ApiMiddlewares is a list of middleware functions that will be applied to all API requests
// this list can be modified to add or remove middlewares
// the order of the middlewares is important, it will be applied in the order they are listed
var ApiMiddlewares = []middleware.Middleware{
	middleware.Security,
	middleware.Trazability,
	middleware.Checks,
}

func Start(configuration configuration.APIConfiguration, endpoints []models.Endpoint) {

	// show log app title and start router
	log.Println(configuration.ApiName)

	// CORS configuration
	// http.HandleFunc("OPTIONS", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	w.Header().Set("Access-Control-Max-Age", "3600")
	// 	w.WriteHeader(http.StatusNoContent)
	// })

	// Add API path to endpoints
	newEndpoints := []models.Endpoint{}
	for _, endpoint := range endpoints {
		endpoint.Path = API_PATH + configuration.ApiName + "/" + configuration.Version + "/" + endpoint.Path
		newEndpoints = append(newEndpoints, endpoint)
	}

	// Register endpoints
	registerEndpoints(newEndpoints)

	// Start listening HTTP requests
	log.Printf("API started on http://%s:%s%s", configuration.Ip, configuration.Port, API_PATH)
	state := http.ListenAndServe(configuration.Ip+":"+configuration.Port, nil)
	log.Print(state.Error())

}

func registerEndpoints(endpoints []models.Endpoint) {

	for _, endpoint := range endpoints {

		log.Printf("Endpoint %s registered. \n", endpoint.Path)

		switch endpoint.Method {
		case models.GetMethod:
			endpoint.Path = "GET " + endpoint.Path
		case models.PostMethod:
			endpoint.Path = "POST " + endpoint.Path
		case models.PutMethod:
			endpoint.Path = "PUT " + endpoint.Path
		case models.DeleteMethod:
			endpoint.Path = "DELETE " + endpoint.Path
		case models.PatchMethod:
			endpoint.Path = "PATCH " + endpoint.Path
		}

		http.HandleFunc(endpoint.Path, func(w http.ResponseWriter, r *http.Request) {

			// log the request
			log.Printf("%s", endpoint.Path)

			// enable CORS
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "3600")

			// create basic api context
			context := &models.ApiContext{
				Trazability: models.Trazability{
					Endpoint: endpoint,
				},
			}

			// Get request data
			err := middleware.Request(r, context)
			if nil != err {
				jsonResponse(w, err.Status, err)
				return
			}

			// Apply middleware to the request
			err = applyMiddleware(context)
			if nil != err {
				jsonResponse(w, err.Status, err)
				return
			}

			// Execute the endpoint
			err = middleware.Response(context)
			if nil != err {
				jsonResponse(w, err.Status, err)
				return
			}

			// Send response
			jsonResponse(w, context.Response.Code, context.Response)

		})
	}
}

func applyMiddleware(context *models.ApiContext) *models.Error {

	for _, middleware := range ApiMiddlewares {
		err := middleware(context)
		if nil != err {
			return err
		}
	}

	return nil

}

func jsonResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set(CONTENT_TYPE_HEADER, "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(response)
}
