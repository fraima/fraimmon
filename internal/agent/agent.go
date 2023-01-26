package agent

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"fraima.io/fraimmon/internal/dtype"
	"fraima.io/fraimmon/internal/wrong"
)

func newMetrics(m dtype.Metrics) {
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

	m.Gauges[0] = dtype.Gauge{
		Name:  "TotalAlloc",
		Value: float64(memStat.TotalAlloc),
	}
	m.Gauges[1] = dtype.Gauge{
		Name:  "Alloc",
		Value: float64(memStat.Alloc),
	}
	m.Gauges[2] = dtype.Gauge{
		Name:  "BuckHashSys",
		Value: float64(memStat.BuckHashSys),
	}
	m.Gauges[3] = dtype.Gauge{
		Name:  "Frees",
		Value: float64(memStat.Frees),
	}
	m.Gauges[4] = dtype.Gauge{
		Name:  "GCCPUFraction",
		Value: float64(memStat.GCCPUFraction),
	}
	m.Gauges[5] = dtype.Gauge{
		Name:  "GCSys",
		Value: float64(memStat.GCSys),
	}
	m.Gauges[6] = dtype.Gauge{
		Name:  "HeapAlloc",
		Value: float64(memStat.HeapAlloc),
	}
	m.Gauges[7] = dtype.Gauge{
		Name:  "HeapIdle",
		Value: float64(memStat.HeapIdle),
	}
	m.Gauges[8] = dtype.Gauge{
		Name:  "HeapInuse",
		Value: float64(memStat.HeapInuse),
	}
	m.Gauges[9] = dtype.Gauge{
		Name:  "HeapObjects",
		Value: float64(memStat.HeapObjects),
	}
	m.Gauges[10] = dtype.Gauge{
		Name:  "HeapReleased",
		Value: float64(memStat.HeapReleased),
	}
	m.Gauges[11] = dtype.Gauge{
		Name:  "HeapSys",
		Value: float64(memStat.HeapSys),
	}
	m.Gauges[12] = dtype.Gauge{
		Name:  "LastGC",
		Value: float64(memStat.LastGC),
	}
	m.Gauges[12] = dtype.Gauge{
		Name:  "Lookups",
		Value: float64(memStat.Lookups),
	}
	m.Gauges[13] = dtype.Gauge{
		Name:  "TotalAlloc",
		Value: float64(memStat.TotalAlloc),
	}
	m.Gauges[14] = dtype.Gauge{
		Name:  "MCacheInuse",
		Value: float64(memStat.MCacheInuse),
	}
	m.Gauges[15] = dtype.Gauge{
		Name:  "MCacheSys",
		Value: float64(memStat.MCacheSys),
	}
	m.Gauges[16] = dtype.Gauge{
		Name:  "MSpanInuse",
		Value: float64(memStat.MSpanInuse),
	}
	m.Gauges[17] = dtype.Gauge{
		Name:  "MSpanSys",
		Value: float64(memStat.MSpanSys),
	}
	m.Gauges[18] = dtype.Gauge{
		Name:  "Mallocs",
		Value: float64(memStat.Mallocs),
	}
	m.Gauges[19] = dtype.Gauge{
		Name:  "NextGC",
		Value: float64(memStat.NextGC),
	}
	m.Gauges[20] = dtype.Gauge{
		Name:  "NumForcedGC",
		Value: float64(memStat.NumForcedGC),
	}
	m.Gauges[21] = dtype.Gauge{
		Name:  "NumGC",
		Value: float64(memStat.NumGC),
	}
	m.Gauges[22] = dtype.Gauge{
		Name:  "OtherSys",
		Value: float64(memStat.OtherSys),
	}
	m.Gauges[23] = dtype.Gauge{
		Name:  "PauseTotalNs",
		Value: float64(memStat.PauseTotalNs),
	}
	m.Gauges[24] = dtype.Gauge{
		Name:  "StackInuse",
		Value: float64(memStat.StackInuse),
	}
	m.Gauges[25] = dtype.Gauge{
		Name:  "StackSys",
		Value: float64(memStat.StackSys),
	}
	m.Gauges[26] = dtype.Gauge{
		Name:  "Sys",
		Value: float64(memStat.Sys),
	}
	m.Gauges[27] = dtype.Gauge{
		Name:  "RandomValue",
		Value: rand.Float64(),
	}
	m.Counters[0] = dtype.Counter{
		Name:  "PollCount",
		Value: pollCount,
	}

}

// test+
func urlTreatment(baseUrl string, item interface{}, metricType string) (string, error) {

	var name, value string

	switch i := item.(type) {

	case dtype.Gauge:
		name = i.Name
		value = strconv.FormatFloat(i.Value, 'f', -1, 64)

	case dtype.Counter:
		name = i.Name
		value = fmt.Sprint(i.Value)

	default:
		return "", wrong.ErrNotFound
	}

	url := baseUrl + "/update/" + metricType + "/" + name + "/" + value

	// TODO валидация урлы вопрос ментору

	return url, nil
}

// test+
func metricTreatment(baseUrl string, metrics dtype.Metrics) ([]string, error) {

	var s []string

	for _, item := range metrics.Gauges {
		url, err := urlTreatment(baseUrl, item, "gauge")

		if err != nil {
			return s, err
		}

		s = append(s, url)
	}

	for _, item := range metrics.Counters {
		url, err := urlTreatment(baseUrl, item, "counter")

		if err != nil {
			return s, err
		}

		s = append(s, url)
	}

	return s, nil
}

func pusher(urlList []string) error {

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

func NewPusher(pushInterval int, mainUrl string, m dtype.Metrics) error {

	newTimerPush := time.Duration(pushInterval) * time.Second
	for {
		<-time.After(newTimerPush)

		urlList, err := metricTreatment(mainUrl, m)
		if err != nil {
			return err
		}

		err = pusher(urlList)
		if err != nil {
			return err
		}
	}
}

func NewScraper(pollInterval int, m dtype.Metrics) {

	newTimerPoll := time.Duration(pollInterval) * time.Second
	for {
		<-time.After(newTimerPoll)
		newMetrics(m)
	}

}
