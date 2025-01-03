package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type APIConfiguration struct {
	Ip      string
	Port    string
	Version string
	ApiName string

	CorsAccessControlAllowOrigin  string
	CorsAccessControlAllowMethods string
	CorsAccessControlAllowHeaders string
	CorsAccessControlMaxAge       string
}

// Load configuration from given env file
func LoadConfiguration(path string) APIConfiguration {

	err := godotenv.Load(path)

	if nil != err {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	var configuration = APIConfiguration{
		Ip:      os.Getenv("IP"),
		Port:    os.Getenv("PORT"),
		Version: os.Getenv("VERSION"),
		ApiName: os.Getenv("API_NAME"),

		CorsAccessControlAllowOrigin:  os.Getenv("CORS_ORIGIN"),
		CorsAccessControlAllowMethods: os.Getenv("CORS_METHODS"),
		CorsAccessControlAllowHeaders: os.Getenv("CORS_HEADERS"),
		CorsAccessControlMaxAge:       os.Getenv("CORS_MAX_AGE"),
	}

	setDefaultVariablesIfNeeded(&configuration)
	printCurrentApiConfiguration(&configuration)
	return configuration
}

// Set defualt variables to configuration if needed
func setDefaultVariablesIfNeeded(configuration *APIConfiguration) {

	if "" == configuration.Ip {
		configuration.Ip = "127.0.0.1"
	}

	if "" == configuration.Port {
		configuration.Ip = "8080"
	}

	if "" == configuration.ApiName {
		configuration.Ip = "api"
	}

	if "" == configuration.Version {
		configuration.Ip = "v1"
	}

	if "" == configuration.Version {
		configuration.Ip = "v1"
	}

}

// Print currentApiConfiguration
func printCurrentApiConfiguration(configuration *APIConfiguration) {
	log.Println()
	log.Println("Configuration variables")
	log.Println()

	log.Printf("IP: %s\n", configuration.Ip)
	log.Printf("PORT: %s\n", configuration.Port)
	log.Printf("VERSION: %s\n", configuration.Version)
	log.Printf("API_NAME: %s\n", configuration.ApiName)
	log.Println()
	log.Printf("CORS_ORIGIN: %s", configuration.CorsAccessControlAllowOrigin)
	log.Printf("CORS_METHODS: %s", configuration.CorsAccessControlAllowMethods)
	log.Printf("CORS_HEADERS: %s", configuration.CorsAccessControlAllowHeaders)
	log.Printf("CORS_MAX_AGE: %s", configuration.CorsAccessControlMaxAge)
}

// Get if current environment is development
func (APIConfiguration) IsDevelopment() bool {
	return os.Getenv("ENV") == "development"
}
