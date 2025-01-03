package main

import (
	apicommon "github.com/akrck02/godot-api-template"
	"github.com/akrck02/godot-api-template/configuration"
	"github.com/akrck02/godot-api-template/models"
	"github.com/akrck02/godot-api-template/services"
)

const ENV_FILE_PATH = "../.env"

func main() {

	config := configuration.LoadConfiguration(ENV_FILE_PATH)

	apicommon.Start(
		config,
		[]models.Endpoint{
			{
				Path:     "health",
				Method:   models.GetMethod,
				Listener: services.Health,
				Checks:   services.EmptyCheck,
				Secured:  false,
				Database: true,
			},
		},
	)
}
