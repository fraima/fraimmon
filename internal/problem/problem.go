package problem

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

func StorageErrToStatus(err error) int {
	switch err {

	case ErrAlreadyExists:
		return http.StatusConflict

	case ErrNotFound:
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
