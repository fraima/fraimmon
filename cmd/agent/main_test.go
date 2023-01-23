package main

import (
	"reflect"
	"testing"
)

func TestPusher(t *testing.T) {
	type args struct {
		urlList []string
	}
	tests := []struct {
		name    string
		fields  args
		wantErr bool
	}{
		{
			name: "test success",
			fields: args{
				urlList: []string{"http://localhost:8080/update/gauge/totalalloc/215064"},
			},
			wantErr: false,
		},
		{
			name: "test failed <connection refused>",
			fields: args{
				urlList: []string{"http://localhost:8081/update/gauge/totalalloc/215064"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Pusher(tt.fields.urlList); (err != nil) != tt.wantErr {
				t.Errorf("Pusher() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestUrlTreatment(t *testing.T) {

	type args struct {
		baseUrl string
		metrics Metrics
	}
	mm := Metrics{
		Gauges:   make([]Gauge, 2),
		Counters: make([]Counter, 1),
	}
	mm.Gauges[0] = Gauge{
		Name:  "Alloc",
		Value: 2,
	}
	mm.Gauges[1] = Gauge{
		Name:  "TotalAlloc",
		Value: 0.111000111,
	}
	mm.Counters[0] = Counter{
		Name:  "PollCount",
		Value: 111,
	}

	tests := []struct {
		name   string
		fields args
		want   []string
	}{
		{
			name: "test success 1",
			want: []string{
				"http://localhost:8080/update/gauge/alloc/2",
				"http://localhost:8080/update/gauge/totalalloc/0.111000111",
				"http://localhost:8080/update/counter/pollcount/111",
			},
			fields: args{
				baseUrl: "http://localhost:8080",
				metrics: mm,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UrlTreatment(tt.fields.baseUrl, mm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("urlTreatment() = %v, want %v", got, tt.want)
			}
		})
	}
}
