package config

import "os"

// Return the status of an environment variable, and log message if needed
func GetBooleanConfiguration(environment string) bool {
	if os.Getenv(environment) == "true" {
		return true
	} else {

		return false
	}
}

func GetStringConfiguration(environment string) string {
	return os.Getenv(environment)
}
