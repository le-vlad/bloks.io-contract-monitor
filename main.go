package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	ramQuota = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "eos_ram_quota",
		Help: "The total amount of RAM",
	})

	ramUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "eos_ram_used",
		Help: "The total amount of used RAM",
	})

	cpuUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "eos_cpu_used",
		Help: "The total amount of used CPU",
	})

	cpuQuota = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "eos_cpu_quota",
		Help: "The total amount of CPU",
	})

	netUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "eos_net_used",
		Help: "The total amount of used CPU",
	})

	netQuota = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "eos_net_quota",
		Help: "The total amount of CPU",
	})
)

func main() {
	contractName := os.Getenv("CONTRACT_NAME")
	c := make(chan int)

	go startPrometheus()
	go heartBeat(c)
	for range c {
		if err := fetchState(contractName); err != nil {
			panic(err)
		}
	}
}

func startPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func fetchState(contract string) (err error) {
	var contractState Account
	if contractState, err = stateFetcher(contract); err != nil {
		return err
	}

	ramQuota.Set(float64(contractState.RAMQuota))
	ramQuota.Write(&dto.Metric{})

	ramUsed.Set(float64(contractState.RAMUsage))
	ramUsed.Write(&dto.Metric{})

	cpuUsed.Set(float64(contractState.CPULimit.Used))
	cpuUsed.Write(&dto.Metric{})

	cpuQuota.Set(float64(contractState.CPULimit.Available))
	cpuQuota.Write(&dto.Metric{})

	netUsed.Set(float64(contractState.NetLimit.Used))
	netUsed.Write(&dto.Metric{})

	netQuota.Set(float64(contractState.NetLimit.Available))
	netQuota.Write(&dto.Metric{})

	fmt.Println("UPDATE METRIC")

	return
}

func heartBeat(beatChan chan int) {
	var (
		intervalInSec int
		err           error
	)
	if intervalInSec, err = strconv.Atoi(os.Getenv("BEAT_INTERVAL")); err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Duration(intervalInSec) * time.Second)
	for range ticker.C {
		beatChan <- 1
	}
}
