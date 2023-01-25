package server

import (
	"reflect"
	"testing"

	"fraima.io/fraimmon/internal/types"
)

func TestUrlTreatment(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name  string
		args  args
		want  types.MetricItem
		want1 int
	}{
		{
			name: "gauge success",
			args: args{
				uri: "/update/gauge/RandomValue/0.4377141871869802",
			},
			want: types.MetricItem{
				Name:  "RandomValue",
				Value: "0.4377141871869802",
				Type:  "gauge",
			},
			want1: 200,
		},
		{
			name: "counter success",
			args: args{
				uri: "/update/counter/PollCount/6443",
			},
			want: types.MetricItem{
				Name:  "PollCount",
				Value: "6443",
				Type:  "counter",
			},
			want1: 200,
		},
		{
			name: "counter failed with incorrect type Value in POST",
			args: args{
				uri: "/update/counter/testCounter/none",
			},
			want: types.MetricItem{
				Name:  "testCounter",
				Value: "none",
				Type:  "counter",
			},
			want1: 200,
		},
		{
			name: "gauge failed with incorrect type Value in POST",
			args: args{
				uri: "/update/gauge/testGauge/none",
			},
			want: types.MetricItem{
				Name:  "testGauge",
				Value: "none",
				Type:  "gauge",
			},
			want1: 200,
		},
		{
			name: "type metric is failed in POST",
			args: args{
				uri: "/update/none/testGauge/100",
			},
			want:  types.MetricItem{},
			want1: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := UrlTreatment(tt.args.uri)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UrlTreatment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UrlTreatment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
