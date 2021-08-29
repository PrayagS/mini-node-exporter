package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/PrayagS/mini-node-exporter/pkg/procfs"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check requested...")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 OK"))
}

func GetHostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := procfs.Hostname()
	if err != nil {
		log.Panicln("error retrieving hostname")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hostname))
}

func GetUptime(w http.ResponseWriter, r *http.Request) {
	systemUptime, err := procfs.Uptime()
	if err != nil {
		log.Panicln("error retrieving uptime")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatFloat(systemUptime, 'f', -1, 64)))
}

func GetLoadAvg(w http.ResponseWriter, r *http.Request) {
	loadavg, err := procfs.AvgLoad()
	if err != nil {
		log.Panicln("error retrieving average load")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loadavg)
}
