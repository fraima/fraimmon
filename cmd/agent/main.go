package main

import (
	"fmt"

	"fraima.io/fraimmon/internal/agent"
	"fraima.io/fraimmon/internal/types"
)

func main() {

	var pollInterval int
	var pushInterval int

	m := types.Metrics{
		Gauges:   make([]types.Gauge, 28),
		Counters: make([]types.Counter, 1),
	}

	pollInterval = 1
	pushInterval = 2

	mainUrl := "http://localhost:8080"

	go agent.NewScraper(pollInterval, m)

	err := agent.NewPusher(pushInterval, mainUrl, m)
	if err != nil {
		fmt.Printf("NewPusher failed with error: %s", err)
	}

}
