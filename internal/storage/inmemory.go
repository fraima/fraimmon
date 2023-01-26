package storage

import (
	// "log"

	"net/http"
	"sync"

	"fraima.io/fraimmon/internal/dtype"
)

type InMemory struct {
	lock sync.Mutex
	g    map[string]dtype.Gauge
	c    map[string]dtype.Counter
}

func NewInMemory() *InMemory {
	return &InMemory{
		g: make(map[string]dtype.Gauge),
		c: make(map[string]dtype.Counter),
	}
}

func (s *InMemory) Get(m interface{}) (interface{}, int) {

	s.lock.Lock()
	defer s.lock.Unlock()

	switch i := m.(type) {

	case dtype.Counter:
		if v, ok := s.c[i.Name]; ok {
			return v.Value, http.StatusOK
		}

	case dtype.Gauge:
		if v, ok := s.g[i.Name]; ok {
			return v.Value, http.StatusOK
		}

	default:
		return nil, http.StatusNotFound

	}

	return nil, http.StatusNotFound
}

func (s *InMemory) Put(m interface{}) int {

	s.lock.Lock()
	defer s.lock.Unlock()

	switch i := m.(type) {

	case dtype.Counter:
		var currentValue, newValue int64

		currentValue = s.c[i.Name].Value
		newValue = currentValue + i.Value
		i.Value = newValue

		s.c[i.Name] = i
		return http.StatusOK

	case dtype.Gauge:
		s.g[i.Name] = i
		return http.StatusOK

	default:
		return http.StatusNotFound
	}

}
