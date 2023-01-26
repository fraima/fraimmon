package util

import (
	"reflect"
	"testing"

	"fraima.io/fraimmon/internal/dtype"
)

func TestURLTreatment(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 int
	}{
		{
			name: "gauge success",
			args: args{
				uri: "/update/gauge/RandomValue/0.4377141871869802",
			},
			want: dtype.Gauge{
				Name:  "RandomValue",
				Value: 0.4377141871869802,
			},
			want1: 200,
		},
		{
			name: "counter success",
			args: args{
				uri: "/update/counter/PollCount/6443",
			},
			want: dtype.Counter{
				Name:  "PollCount",
				Value: 6443,
			},
			want1: 200,
		},
		{
			name: "counter failed with incorrect type Value in POST",
			args: args{
				uri: "/update/counter/testCounter/none",
			},
			want:  dtype.Counter{},
			want1: 400,
		},
		{
			name: "gauge failed with incorrect value in POST",
			args: args{
				uri: "/update/gauge/testGauge/none",
			},
			want:  dtype.Gauge{},
			want1: 400,
		},
		{
			name: "gauge failed with incorrect url in POST",
			args: args{
				uri: "/update/gauge/",
			},
			want:  nil,
			want1: 404,
		},
		{
			name: "counter failed with incorrect url in POST",
			args: args{
				uri: "/update/counter/",
			},
			want:  nil,
			want1: 404,
		},
		{
			name: "unknown failed with incorrect type in POST",
			args: args{
				uri: "/update/unknown/",
			},
			want:  nil,
			want1: 501,
		},
		{
			name: "unknown failed with incorrect type in POST",
			args: args{
				uri: "/update/unknown/testCounter/100",
			},
			want:  nil,
			want1: 501,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := URLTreatment(tt.args.uri)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLTreatment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("URLTreatment() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
