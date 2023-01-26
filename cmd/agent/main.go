package main

import (
	"fmt"

	"fraima.io/fraimmon/internal/agent"
	"fraima.io/fraimmon/internal/dtype"
)

func main() {

	var pollInterval int
	var pushInterval int

	m := dtype.Metrics{
		Gauges:   make([]dtype.Gauge, 28),
		Counters: make([]dtype.Counter, 1),
	}

	pollInterval = 1
	pushInterval = 10

	mainUrl := "http://localhost:8080"

	go agent.NewScraper(pollInterval, m)

	err := agent.NewPusher(pushInterval, mainUrl, m)
	if err != nil {
		fmt.Printf("NewPusher failed with error: %s", err)
	}

}
