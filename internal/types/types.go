package types

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
