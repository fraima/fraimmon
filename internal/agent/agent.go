package agent

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"fraima.io/fraimmon/internal/types"
)

func NewMetrics(m types.Metrics) {
	var memStat runtime.MemStats
	var pollCount int64

	for _, item := range m.Counters {

		if item.Name == "PollCount" {
			pollCount = item.Value + 1
		} else {
			pollCount = 1
		}

	}
	runtime.ReadMemStats(&memStat)

	m.Gauges[0] = types.Gauge{
		Name:  "TotalAlloc",
		Value: float64(memStat.TotalAlloc),
	}
	m.Gauges[1] = types.Gauge{
		Name:  "Alloc",
		Value: float64(memStat.Alloc),
	}
	m.Gauges[2] = types.Gauge{
		Name:  "BuckHashSys",
		Value: float64(memStat.BuckHashSys),
	}
	m.Gauges[3] = types.Gauge{
		Name:  "Frees",
		Value: float64(memStat.Frees),
	}
	m.Gauges[4] = types.Gauge{
		Name:  "GCCPUFraction",
		Value: float64(memStat.GCCPUFraction),
	}
	m.Gauges[5] = types.Gauge{
		Name:  "GCSys",
		Value: float64(memStat.GCSys),
	}
	m.Gauges[6] = types.Gauge{
		Name:  "HeapAlloc",
		Value: float64(memStat.HeapAlloc),
	}
	m.Gauges[7] = types.Gauge{
		Name:  "HeapIdle",
		Value: float64(memStat.HeapIdle),
	}
	m.Gauges[8] = types.Gauge{
		Name:  "HeapInuse",
		Value: float64(memStat.HeapInuse),
	}
	m.Gauges[9] = types.Gauge{
		Name:  "HeapObjects",
		Value: float64(memStat.HeapObjects),
	}
	m.Gauges[10] = types.Gauge{
		Name:  "HeapReleased",
		Value: float64(memStat.HeapReleased),
	}
	m.Gauges[11] = types.Gauge{
		Name:  "HeapSys",
		Value: float64(memStat.HeapSys),
	}
	m.Gauges[12] = types.Gauge{
		Name:  "LastGC",
		Value: float64(memStat.LastGC),
	}
	m.Gauges[12] = types.Gauge{
		Name:  "Lookups",
		Value: float64(memStat.Lookups),
	}
	m.Gauges[13] = types.Gauge{
		Name:  "TotalAlloc",
		Value: float64(memStat.TotalAlloc),
	}
	m.Gauges[14] = types.Gauge{
		Name:  "MCacheInuse",
		Value: float64(memStat.MCacheInuse),
	}
	m.Gauges[15] = types.Gauge{
		Name:  "MCacheSys",
		Value: float64(memStat.MCacheSys),
	}
	m.Gauges[16] = types.Gauge{
		Name:  "MSpanInuse",
		Value: float64(memStat.MSpanInuse),
	}
	m.Gauges[17] = types.Gauge{
		Name:  "MSpanSys",
		Value: float64(memStat.MSpanSys),
	}
	m.Gauges[18] = types.Gauge{
		Name:  "Mallocs",
		Value: float64(memStat.Mallocs),
	}
	m.Gauges[19] = types.Gauge{
		Name:  "NextGC",
		Value: float64(memStat.NextGC),
	}
	m.Gauges[20] = types.Gauge{
		Name:  "NumForcedGC",
		Value: float64(memStat.NumForcedGC),
	}
	m.Gauges[21] = types.Gauge{
		Name:  "NumGC",
		Value: float64(memStat.NumGC),
	}
	m.Gauges[22] = types.Gauge{
		Name:  "OtherSys",
		Value: float64(memStat.OtherSys),
	}
	m.Gauges[23] = types.Gauge{
		Name:  "PauseTotalNs",
		Value: float64(memStat.PauseTotalNs),
	}
	m.Gauges[24] = types.Gauge{
		Name:  "StackInuse",
		Value: float64(memStat.StackInuse),
	}
	m.Gauges[25] = types.Gauge{
		Name:  "StackSys",
		Value: float64(memStat.StackSys),
	}
	m.Gauges[26] = types.Gauge{
		Name:  "Sys",
		Value: float64(memStat.Sys),
	}
	m.Gauges[27] = types.Gauge{
		Name:  "RandomValue",
		Value: rand.Float64(),
	}
	m.Counters[0] = types.Counter{
		Name:  "PollCount",
		Value: pollCount,
	}
	a, _ := json.Marshal(m)
	fmt.Println(string(a))
}

func urlTreatment(baseUrl string, item interface{}, metricType string) string {

	var name, value string

	switch i := item.(type) {
	case types.Gauge:
		name = i.Name
		value = strconv.FormatFloat(i.Value, 'f', -1, 64)
	case types.Counter:
		name = i.Name
		value = fmt.Sprint(i.Value)
	}

	url := baseUrl + "/update/" + metricType + "/" + name + "/" + value

	return url
}

// test+
func MetricTreatment(baseUrl string, metrics types.Metrics) []string {

	var s []string

	for _, item := range metrics.Gauges {
		url := urlTreatment(baseUrl, item, "gauge")
		s = append(s, url)
	}

	for _, item := range metrics.Counters {
		url := urlTreatment(baseUrl, item, "counter")
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

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("warning: status code != <%v>", resp.StatusCode)
		}

		defer resp.Body.Close()

	}
	return nil
}

func NewPusher(pushInterval int, mainUrl string, m types.Metrics) error {

	newTimerPush := time.Duration(pushInterval) * time.Second
	for {
		<-time.After(newTimerPush)
		err := Pusher(MetricTreatment(mainUrl, m))
		if err != nil {
			return err
		}
	}
}

func NewScraper(pollInterval int, m types.Metrics) {

	newTimerPoll := time.Duration(pollInterval) * time.Second
	for {
		<-time.After(newTimerPoll)
		NewMetrics(m)
	}

}
