package collectors

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector which collects the system uptime
type uptimeCollector struct {
	metric    *prometheus.Desc
	valueType prometheus.ValueType
}

// Collector which collects the average load values
type loadAvgCollector struct {
	metric    *prometheus.Desc
	valueType prometheus.ValueType
}

// Create a new prometheus registry
var Registry = prometheus.NewRegistry()

// Define the prometheus metrics
var (
	GaugeUptime = prometheus.NewDesc(
		"node_uptime",
		"No. of seconds the system is running",
		[]string{"hostname"},
		nil,
	)
	GaugeLoadAvg = prometheus.NewDesc(
		"node_load",
		"Average load of the system at different intervals",
		[]string{"hostname", "duration"},
		nil,
	)
)

// Function which registers the newly created collectors to the registry
func RegisterCollectors() {

	uc, err := NewUptimeCollector()
	if err != nil {
		log.Println("error registering uptime collector")
	}

	lc, err := NewLoadAvgCollector()
	if err != nil {
		log.Println("error registering average load collector")
	}

	Registry.MustRegister(uc)
	Registry.MustRegister(lc)
}
