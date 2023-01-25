package storage

import (
	"fraima.io/fraimmon/internal/types"
)

type Storage interface {
	Get(m types.MetricItem) (interface{}, int)
	Put(m types.MetricItem) int
}
