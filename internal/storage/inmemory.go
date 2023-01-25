package storage

import (
	"strconv"
	"sync"

	"fraima.io/fraimmon/internal/types"
)

type InMemory struct {
	lock sync.Mutex
	g    map[string]types.Gauge
	c    map[string]types.Counter
}

func NewInMemory() *InMemory {
	return &InMemory{
		g: make(map[string]types.Gauge),
		c: make(map[string]types.Counter),
	}
}

func (s *InMemory) Get(key string, metricType string) (interface{}, error) {

	var i interface{}
	s.lock.Lock()
	defer s.lock.Unlock()

	switch metricType {
	case "couneter":
		if v, ok := s.c[key]; ok {
			return v.Value, nil
		}
	case "gauge":
		if v, ok := s.g[key]; ok {
			return v.Value, nil
		}
	}

	return i, ErrNotFound
}

func (s *InMemory) Put(key string, value string, metricType string) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	switch metricType {
	case "couneter":
		var i types.Counter

		v, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return err
		}

		i.Name = key
		i.Value = v

		s.c[key] = i

	case "gauge":
		var i types.Gauge

		v, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return err
		}
		i.Name = key
		i.Value = v

		s.g[key] = i
	}

	return nil
}
