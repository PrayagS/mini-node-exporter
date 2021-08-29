package collectors

import (
	"log"

	"github.com/PrayagS/mini-node-exporter/pkg/procfs"
	"github.com/prometheus/client_golang/prometheus"
)

// NewLoadavgCollector returns a new Collector which collects the load average stats.
func NewLoadAvgCollector() (*loadAvgCollector, error) {
	return &loadAvgCollector{
		metric:    GaugeLoadAvg,
		valueType: prometheus.GaugeValue,
	}, nil
}

// Describe implements the prometheus.Collector interface.
func (c *loadAvgCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metric
}

// Collect implements the prometheus.Collector interface.
func (c *loadAvgCollector) Collect(ch chan<- prometheus.Metric) {
	var hostname string
	var err error
	hostname, err = procfs.Hostname()
	if err != nil {
		// Default the hostname to 'localhost' in case of any error
		hostname = "localhost"
		log.Println("error retrieving hostname. defaulting to localhost")
		return
	}

	loadValues, err := procfs.AvgLoad()

	// See: https://prometheus.io/docs/instrumenting/writing_exporters/#collectors
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			c.metric, c.valueType, 0, hostname,
		)
		ch <- prometheus.MustNewConstMetric(
			c.metric, c.valueType, 0, hostname,
		)
		ch <- prometheus.MustNewConstMetric(
			c.metric, c.valueType, 0, hostname,
		)
		log.Println("error retrieving uptime")
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.metric, c.valueType, loadValues.Load1m, hostname, "1m",
	)
	ch <- prometheus.MustNewConstMetric(
		c.metric, c.valueType, loadValues.Load5m, hostname, "5m",
	)
	ch <- prometheus.MustNewConstMetric(
		c.metric, c.valueType, loadValues.Load15m, hostname, "15m",
	)
}
