package config

import (
	"log"
	"os"
)

var GitApiKey = ""
var GitProvider = ""

func Init() {

	// Checking that git config api key is set
	GitApiKey = getStringConfiguration("CONFIG_GIT_API_KEY")
	if GitApiKey == "" {
		log.Fatal("Expected environment variable CONFIG_GIT_API_KEY, please configure it")
	}
	// Configure git provider
	GitProvider = getStringConfiguration("CONFIG_GIT_PROVIDER")
	switch GitProvider {
	case "github":
		log.Printf("Using github")
	default:
		log.Fatal("Please, configure a valid git provider")
	}
	// verifying that api key is valid

}

// Return the status of an environment variable, and log message if needed
func GetBooleanConfiguration(environment string) bool {
	if os.Getenv(environment) == "true" {
		return true
	} else {

		return false
	}
}

func getStringConfiguration(environment string) string {
	return os.Getenv(environment)
}
