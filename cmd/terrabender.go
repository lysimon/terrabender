package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lysimon/terrabender/pkg/config"
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
	if config.GetBooleanConfiguration("CONFIG_DAEMON_ENABLED") {
		log.Print("Running as a daemon agent, program will not stop")
		// starting daemon for verify, rebase, merge, apply
		webserver()

	} else {
		log.Print("Daemon disabled, only going to the required job")

	}
}

func webserver() {
	// Start webserver only if CONFIG_WEBSERVER_ENABLED is true
	if config.GetBooleanConfiguration("CONFIG_WEBSERVER_ENABLED") {
		log.Print("terrabender.webserver start")

		router := mux.NewRouter()
		router.HandleFunc("/status", status.Status)

		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Print("Webserver disabled")
	}
}

func plan() {

}

func apply() {

}

func rebase() {

}
