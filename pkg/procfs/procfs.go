package procfs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type LoadAvg struct {
	Load1m  float64 `json:"1m"`
	Load5m  float64 `json:"5m"`
	Load15m float64 `json:"15m"`
}

func Hostname() (hostname string, err error) {
	filepath := "/proc/sys/kernel/hostname"

	hostnameInBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading %v: %w", filepath, err)
	}

	if hostnameInBytes[len(hostnameInBytes)-1] == '\n' {
		return string(hostnameInBytes[:len(hostnameInBytes)-1]), nil
	}

	return string(hostnameInBytes), err
}

func Uptime() (systemUptime float64, err error) {
	filepath := "/proc/uptime"

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return 0, err
	}

	uptime := strings.Fields(string(data))
	systemUptime, err = strconv.ParseFloat(uptime[0], 64)
	if err != nil {
		return 0, err
	}

	return systemUptime, err
}

func AvgLoad() (*LoadAvg, error) {
	filepath := "/proc/loadavg"

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	loadavgValues := strings.Fields(string(data))
	if len(loadavgValues) < 3 {
		return nil, errors.New("error parsing average load")
	}
	loads := make([]float64, 3)
	for i, load := range loadavgValues[0:3] {
		loads[i], err = strconv.ParseFloat(load, 64)
		if err != nil {
			return nil, err
		}
	}

	return &LoadAvg{
		Load1m:  loads[0],
		Load5m:  loads[1],
		Load15m: loads[2],
	}, err
}
