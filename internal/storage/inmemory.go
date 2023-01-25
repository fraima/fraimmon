package storage

import (
	// "log"
	"net/http"
	"strconv"
	"sync"

	"fraima.io/fraimmon/internal/problem"
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

func (s *InMemory) Get(m types.MetricItem) (interface{}, int) {

	// log.Printf("инициализация запроса GET inMemory")

	s.lock.Lock()
	defer s.lock.Unlock()

	switch m.Type {
	case "counter":
		if v, ok := s.c[m.Name]; ok {
			return v.Value, http.StatusOK
		}
	case "gauge":
		if v, ok := s.g[m.Name]; ok {
			return v.Value, http.StatusOK
		}
	default:
		return nil, problem.StorageErrToStatus(problem.ErrNotFound)

	}

	return nil, http.StatusOK
}

func (s *InMemory) Put(m types.MetricItem) int {

	// log.Printf("<Put:inMemory> start func")
	// log.Printf("<Put:InMemory> payload <- <Put:server>: %s", m)
	s.lock.Lock()
	defer s.lock.Unlock()

	switch m.Type {
	case "counter":
		// log.Printf("<Put:inMemory> %s", m.Type)
		// log.Printf("<Put:inMemory> goto case counter")
		var i types.Counter

		v, err := strconv.ParseInt(m.Value, 10, 64)

		if err != nil {
			return http.StatusBadRequest
		}

		i.Name = m.Name
		i.Value = v

		s.c[m.Name] = i

	case "gauge":
		// log.Printf("<Put:inMemory> %s", m.Type)
		// log.Printf("<Put:inMemory> goto case gauge")
		var i types.Gauge

		v, err := strconv.ParseFloat(m.Value, 64)

		if err != nil {
			return http.StatusBadRequest
		}

		i.Name = m.Name
		i.Value = v

		s.g[m.Name] = i

	default:
		// log.Printf("<Put:inMemory> %s", m.Type)
		// log.Printf("<Put:inMemory> goto case DEFAULT")
		return problem.StorageErrToStatus(problem.ErrNotFound)
	}
	// log.Printf("<Put:inMemory> exit func")
	return http.StatusOK
}
