package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lysimon/terrabender/pkg/config"
	"github.com/lysimon/terrabender/pkg/git"
	"github.com/lysimon/terrabender/pkg/status"
)

func main() {
	// Initialization, creating s3 buckets with retention policy for the future

	// Start clean up cron job => should be done by tail
	//c := cron.New()
	//c.AddFunc("5 * * * * *", func() { paw.Cron() })
	//c.Start()
	//defer c.Stop()

	config.Init()

	// Start webserver if dae
	if config.Daemon {
		log.Print("Running as a daemon agent, program will not stop")
		// starting daemon for verify, rebase, merge, apply
		webserver()

	} else {
		log.Print("Daemon disabled, only going to run the required jobs")
		// Running all job was something still need to be done
		finished := false
		for !finished {
			finished = true
			if config.PlanEnabled {
				res, err := plan()
				if err != nil {
					log.Fatal("Unable to perform planning step: " + err.Error())
				}
				if res && !finished {
					finished = true
				}
			}

			if config.ApplyEnabled {
				res, err := apply()
				if err != nil {
					log.Fatal("Unable to perform apply step: " + err.Error())
				}
				if res && !finished {
					finished = true
				}
			}

			if config.MergeEnabled {
				res, err := merge()
				if err != nil {
					log.Fatal("Unable to perform merging step: " + err.Error())
				}
				if res && !finished {
					finished = true
				}
			}
		}
	}
}

func webserver() {
	// Start webserver only if CONFIG_WEBSERVER_ENABLED is true
	if config.WebserverEnabled {
		log.Print("terrabender.webserver start")

		router := mux.NewRouter()
		router.HandleFunc("/status", status.Status)

		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Print("Webserver disabled")
	}
}

func plan() (bool, error) {
	//
	_, err := git.Clone("github.com/lysimon/terrabender", "master", "/tmp/toto")
	if err != nil {
		log.Print("Got unexpected error while doing git clone: " + err.Error())
	}
	// Executing terraform plan

	// Check if no error, if error, write output as error + all the string

	// Delete local path
	return true, nil
}

// true if an action was performed, false otherwise
func apply() (bool, error) {

	return true, nil
}

// true if an action was performed, false otherwise
func merge() (bool, error) {

	return true, nil
}
