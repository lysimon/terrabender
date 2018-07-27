package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lysimon/terrabender/pkg/status"
)

func main() {
	// Initialization, creating s3 buckets with retention policy for the future

	// Start clean up cron job => should be done by tail
	//c := cron.New()
	//c.AddFunc("5 * * * * *", func() { paw.Cron() })
	//c.Start()
	//defer c.Stop()

	// Start webserver if dae
	webserver()
}

func webserver() {
	// Start webserver only if CONFIG_WEBSERVER_ENABLED is true
	if os.Getenv("CONFIG_WEBSERVER_ENABLED") == "true" {
		log.Print("terrabender.webserver start")
		router := mux.NewRouter()
		router.HandleFunc("/status", status.Status)

		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Print("environment variable CONFIG_WEBSERVER_ENABLED is not true, webserver is disabled")
		log.Print(os.Getenv("CONFIG_WEBSERVER_ENABLED"))
	}

}
