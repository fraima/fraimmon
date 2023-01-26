package agent

import (
	"reflect"
	"testing"

	"fraima.io/fraimmon/internal/dtype"
)

func Test_urlTreatment(t *testing.T) {
	type test struct {
		Name  string
		Value string
	}
	type args struct {
		baseUrl    string
		item       interface{}
		metricType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "gauge success",
			args: args{
				item: dtype.Gauge{
					Name:  "gaugetest",
					Value: 0.6546546,
				},
				metricType: "gauge",
				baseUrl:    "https://fraima.io",
			},
			want:    "https://fraima.io/update/gauge/gaugetest/0.6546546",
			wantErr: false,
		},
		{
			name: "counter success",
			args: args{
				item: dtype.Counter{
					Name:  "countertest",
					Value: 2,
				},
				metricType: "counter",
				baseUrl:    "https://fraima.io",
			},
			want:    "https://fraima.io/update/counter/countertest/2",
			wantErr: false,
		},
		{
			name: "undefined item type",
			args: args{
				item: test{
					Name:  "undefinedtest",
					Value: "testvalue",
				},
				metricType: "treska",
				baseUrl:    "https://fraima.io",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := urlTreatment(tt.args.baseUrl, tt.args.item, tt.args.metricType)
			if (err != nil) != tt.wantErr {
				t.Errorf("urlTreatment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("urlTreatment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metricTreatment(t *testing.T) {
	type args struct {
		baseUrl string
		metrics dtype.Metrics
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test 1 success",
			args: args{
				baseUrl: "https://fraima.io",
				metrics: dtype.Metrics{
					Gauges: []dtype.Gauge{
						{
							Name:  "gaugetest1",
							Value: 0.546857864,
						},
						{
							Name:  "gaugetest2",
							Value: 0.2,
						},
					},
					Counters: []dtype.Counter{
						{
							Name:  "countertest1",
							Value: 1,
						},
						{
							Name:  "countertest2",
							Value: 152,
						},
					},
				},
			},
			want: []string{
				"https://fraima.io/update/gauge/gaugetest1/0.546857864",
				"https://fraima.io/update/gauge/gaugetest2/0.2",
				"https://fraima.io/update/counter/countertest1/1",
				"https://fraima.io/update/counter/countertest2/152",
			},
			wantErr: false,
		},
		{
			name: "test-2 empty metric list",
			args: args{
				baseUrl: "https://fraima.io",
				metrics: dtype.Metrics{
					Gauges:   []dtype.Gauge{},
					Counters: []dtype.Counter{},
				},
			},
			wantErr: false,
		},
		{
			name: "test-3 metric without value",
			args: args{
				baseUrl: "https://fraima.io",
				metrics: dtype.Metrics{
					Gauges: []dtype.Gauge{
						{
							Name: "gaugetest1",
						},
					},
					Counters: []dtype.Counter{
						{
							Name: "countertest1",
						},
					},
				},
			},
			want: []string{
				"https://fraima.io/update/gauge/gaugetest1/0",
				"https://fraima.io/update/counter/countertest1/0",
			},
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				metrics: dtype.Metrics{
					Gauges: []dtype.Gauge{
						{
							Name: "gaugetest1",
						},
					},
					Counters: []dtype.Counter{
						{
							Name: "countertest1",
						},
					},
				},
				baseUrl: "",
			},

			want: []string{
				"/update/gauge/gaugetest1/0",
				"/update/counter/countertest1/0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := metricTreatment(tt.args.baseUrl, tt.args.metrics)
			if (err != nil) != tt.wantErr {
				t.Errorf("metricTreatment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("metricTreatment() = %v, want %v", got, tt.want)
			}
		})
	}
}
