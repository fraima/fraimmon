package storage

import (
	"sync"
	"testing"

	"fraima.io/fraimmon/internal/types"
)

func TestInMemory_Put(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		g    map[string]types.Gauge
		c    map[string]types.Counter
	}
	type args struct {
		m types.MetricItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				g:    tt.fields.g,
				c:    tt.fields.c,
			}
			if got := s.Put(tt.args.m); got != tt.want {
				t.Errorf("InMemory.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
