package main

import (
	"log"
	"net/http"

	"github.com/PrayagS/mini-node-exporter/pkg/collectors"
	"github.com/PrayagS/mini-node-exporter/web/handlers"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// Define the collectors and register them to the prometheus registry
	collectors.RegisterCollectors()

	r := mux.NewRouter()

	// Attach the routes and their respective handler functions
	r.HandleFunc("/healthz/ready", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/info/hostname", handlers.GetHostname).Methods("GET")
	r.HandleFunc("/info/uptime", handlers.GetUptime).Methods("GET")
	r.HandleFunc("/info/load", handlers.GetLoadAvg).Methods("GET")

	// Attach the prometheus handler to /metrics
	r.Path("/metrics").Handler(promhttp.HandlerFor(collectors.Registry, promhttp.HandlerOpts{}))

	// Apply the CORS middleware to our top-level router, with the defaults.
	log.Fatal(http.ListenAndServe("0.0.0.0:23333", gorillaHandlers.CORS()(r)))
}
