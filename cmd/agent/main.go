package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Gauge struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Counter struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type Metrics struct {
	Gauges   []Gauge   `json:"gauges"`
	Counters []Counter `json:"counters"`
}

var m Metrics

func NewMetrics(pollCount int64) {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	m = Metrics{
		Gauges:   make([]Gauge, 28),
		Counters: make([]Counter, 1),
	}

	m.Gauges[0] = Gauge{
		Name:  "TotalAlloc",
		Value: float64(memStat.TotalAlloc),
	}
	m.Gauges[1] = Gauge{
		Name:  "Alloc",
		Value: float64(memStat.Alloc),
	}
	m.Gauges[2] = Gauge{
		Name:  "BuckHashSys",
		Value: float64(memStat.BuckHashSys),
	}
	m.Gauges[3] = Gauge{
		Name:  "Frees",
		Value: float64(memStat.Frees),
	}
	m.Gauges[4] = Gauge{
		Name:  "GCCPUFraction",
		Value: float64(memStat.GCCPUFraction),
	}
	m.Gauges[5] = Gauge{
		Name:  "GCSys",
		Value: float64(memStat.GCSys),
	}
	m.Gauges[6] = Gauge{
		Name:  "HeapAlloc",
		Value: float64(memStat.HeapAlloc),
	}
	m.Gauges[7] = Gauge{
		Name:  "HeapIdle",
		Value: float64(memStat.HeapIdle),
	}
	m.Gauges[8] = Gauge{
		Name:  "HeapInuse",
		Value: float64(memStat.HeapInuse),
	}
	m.Gauges[9] = Gauge{
		Name:  "HeapObjects",
		Value: float64(memStat.HeapObjects),
	}
	m.Gauges[10] = Gauge{
		Name:  "HeapReleased",
		Value: float64(memStat.HeapReleased),
	}
	m.Gauges[11] = Gauge{
		Name:  "HeapSys",
		Value: float64(memStat.HeapSys),
	}
	m.Gauges[12] = Gauge{
		Name:  "LastGC",
		Value: float64(memStat.LastGC),
	}
	m.Gauges[12] = Gauge{
		Name:  "Lookups",
		Value: float64(memStat.Lookups),
	}
	m.Gauges[14] = Gauge{
		Name:  "MCacheInuse",
		Value: float64(memStat.MCacheInuse),
	}
	m.Gauges[15] = Gauge{
		Name:  "MCacheSys",
		Value: float64(memStat.MCacheSys),
	}
	m.Gauges[16] = Gauge{
		Name:  "MSpanInuse",
		Value: float64(memStat.MSpanInuse),
	}
	m.Gauges[17] = Gauge{
		Name:  "MSpanSys",
		Value: float64(memStat.MSpanSys),
	}
	m.Gauges[18] = Gauge{
		Name:  "Mallocs",
		Value: float64(memStat.Mallocs),
	}
	m.Gauges[19] = Gauge{
		Name:  "NextGC",
		Value: float64(memStat.NextGC),
	}
	m.Gauges[20] = Gauge{
		Name:  "NumForcedGC",
		Value: float64(memStat.NumForcedGC),
	}
	m.Gauges[21] = Gauge{
		Name:  "NumGC",
		Value: float64(memStat.NumGC),
	}
	m.Gauges[22] = Gauge{
		Name:  "OtherSys",
		Value: float64(memStat.OtherSys),
	}
	m.Gauges[23] = Gauge{
		Name:  "PauseTotalNs",
		Value: float64(memStat.PauseTotalNs),
	}
	m.Gauges[24] = Gauge{
		Name:  "StackInuse",
		Value: float64(memStat.StackInuse),
	}
	m.Gauges[25] = Gauge{
		Name:  "StackSys",
		Value: float64(memStat.StackSys),
	}
	m.Gauges[26] = Gauge{
		Name:  "Sys",
		Value: float64(memStat.Sys),
	}
	m.Gauges[27] = Gauge{
		Name:  "TotalAlloc",
		Value: float64(memStat.TotalAlloc),
	}
	m.Counters[0] = Counter{
		Name:  "PollCount",
		Value: pollCount,
	}
}

// test+
func UrlTreatment(baseUrl string, metrics Metrics) []string {

	var s []string

	for _, item := range metrics.Gauges {
		gaugeName := strings.ToLower(item.Name)
		gaugeValue := strconv.FormatFloat(item.Value, 'f', -1, 64)
		gaugeType := "gauge"
		url := baseUrl + "/update/" + gaugeType + "/" + gaugeName + "/" + gaugeValue

		s = append(s, url)
	}

	for _, item := range metrics.Counters {
		gaugeName := strings.ToLower(item.Name)
		gaugeValue := fmt.Sprint(item.Value)
		gaugeType := "counter"
		url := baseUrl + "/update/" + gaugeType + "/" + gaugeName + "/" + gaugeValue

		s = append(s, url)
	}

	return s
}

// test+
func Pusher(urlList []string) error {

	for _, url := range urlList {

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("warning: status code != <%v>", resp.StatusCode)
		}

		resp.Body.Close()

	}
	return nil
}

func NewPusher(pushInterval int, mainUrl string) error {

	newTimerPush := time.Duration(pushInterval) * time.Second
	for {
		<-time.After(newTimerPush)

		err := Pusher(UrlTreatment(mainUrl, m))

		if err != nil {
			return err
		}

	}
}

func NewScraper(pollInterval int, ch chan error) {

	var pollCount int64

	newTimerPoll := time.Duration(pollInterval) * time.Second
	for {
		<-time.After(newTimerPoll)

		pollCount += 1
		NewMetrics(pollCount)
	}
	<-ch
}

func main() {

	var chScrape chan error
	var pollInterval int
	var pushInterval int

	pollInterval = 2
	pushInterval = 4

	mainUrl := "http://localhost:8080"

	go NewScraper(pollInterval, chScrape)

	err := NewPusher(pushInterval, mainUrl)
	if err != nil {
		panic(1)
	}

}
