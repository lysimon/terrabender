package config

import (
	"log"
	"os"
)

var GitApiKey = ""
var GitProvider = ""

// Comma separated list of url project
var GitUrls = ""

// Local path to use, usually /tmp
var LocalPath = ""

//
var Daemon = false

var PlanEnabled = true
var ApplyEnabled = true
var MergeEnabled = true

var WebserverEnabled = true

func Init() {

	Daemon = GetBooleanConfiguration("CONFIG_DAEMON_ENABLED")
	PlanEnabled = GetBooleanConfiguration("CONFIG_PLAN_ENABLED")

	// Usually, merge and apply should be done by the same job
	ApplyEnabled = GetBooleanConfiguration("CONFIG_APPLY_ENABLED")
	MergeEnabled = GetBooleanConfiguration("CONFIG_MERGE_ENABLED")

	// Get webserver configuration. Webserver will only be enabled in daemon mode
	WebserverEnabled = GetBooleanConfiguration("CONFIG_WEBSERVER_ENABLED")

	// Local path, default to /tmp. Will be verified as writable on startup
	LocalPath = GetStringConfiguration("CONFIG_LOCAL_PATH")
	// Setting default vaue for local path
	if LocalPath == "" {
		LocalPath = "/tmp"
		log.Print("Expected environment variable CONFIG_LOCAL_PATH, default to /tmp")
	} else {
		// Verify that we can create at least one file in it, and fail if not
	}
	// Checking that git config api key is set
	GitApiKey = GetStringConfiguration("CONFIG_GIT_API_KEY")
	if GitApiKey == "" {
		log.Fatal("Expected environment variable CONFIG_GIT_API_KEY, please configure it")
	}

	// Configure git provider
	GitProvider = GetStringConfiguration("CONFIG_GIT_PROVIDER")
	switch GitProvider {
	case "github":
		log.Printf("Using github")
	default:
		log.Fatal("Please, configure a valid git provider with variable CONFIG_GIT_PROVIDER, got " + GitProvider)
	}
	// Getting list of git repository
	GitUrls = GetStringConfiguration("CONFIG_GIT_URLS")
	if GitUrls == "" {
		log.Fatal("Expected environment variable CONFIG_GIT_URLS, please configure it")
	} else {

	}

}

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
