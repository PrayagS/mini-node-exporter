package collectors

import (
	"log"

	"github.com/PrayagS/mini-node-exporter/pkg/procfs"
	"github.com/prometheus/client_golang/prometheus"
)

// NewUptimeCollector returns a new Collector which collects the system time.
func NewUptimeCollector() (*uptimeCollector, error) {
	return &uptimeCollector{
		metric:    GaugeUptime,
		valueType: prometheus.GaugeValue,
	}, nil
}

// Describe implements the prometheus.Collector interface.
func (c *uptimeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metric
}

// Collect implements the prometheus.Collector interface.
func (c *uptimeCollector) Collect(ch chan<- prometheus.Metric) {
	var hostname string
	var err error
	hostname, err = procfs.Hostname()
	if err != nil {
		// Default the hostname to 'localhost' in case of any error
		hostname = "localhost"
		log.Println("error retrieving hostname. defaulting to localhost")
		return
	}

	uptime_f64, err := procfs.Uptime()

	// See: https://prometheus.io/docs/instrumenting/writing_exporters/#collectors
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			c.metric, c.valueType, 0, hostname,
		)
		log.Println("error retrieving uptime")
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.metric, c.valueType, uptime_f64, hostname,
	)
}
