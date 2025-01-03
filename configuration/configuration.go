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
}

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
	}

	checkCompulsoryVariables(configuration)
	return configuration
}

func checkCompulsoryVariables(Configuration APIConfiguration) {
	log.Println()
	log.Println("Configuration variables")
	log.Println()
	log.Printf("IP: %s\n", Configuration.Ip)
	log.Printf("PORT: %s\n", Configuration.Port)
	log.Printf("VERSION: %s\n", Configuration.Version)
	log.Printf("API_NAME: %s\n", Configuration.ApiName)
}

func (APIConfiguration) IsDevelopment() bool {
	return os.Getenv("ENV") == "development"
}
