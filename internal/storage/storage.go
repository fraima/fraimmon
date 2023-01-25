package storage

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type Storage interface {
	Get(key string, metricType string) (interface{}, error)
	Put(key, value string, metricType string) error
}
