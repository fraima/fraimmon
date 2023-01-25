package server

import (
	"fmt"
	"net/http"
	"regexp"

	"fraima.io/fraimmon/internal/storage"
)

type Server struct {
	storage storage.Storage
}

func UrlTreatmentPush(urli string) (string, string, string) {

	var metricType string
	var metricName string
	var metricValue string

	re := regexp.MustCompile(`\/update\/(counter|gauge)\/(\w*)\/([0-9]*.[0-9]*)`)
	sliceReg := re.FindStringSubmatch(urli)

	metricType = sliceReg[1]
	metricName = sliceReg[2]
	metricValue = sliceReg[3]

	return metricType, metricName, metricValue

}

func UrlTreatmentGet(urli string) (string, string) {
	var metricType string
	var metricName string

	re := regexp.MustCompile(`\/update\/(counter|gauge)\/(\w*)`)
	sliceReg := re.FindStringSubmatch(urli)

	metricType = sliceReg[1]
	metricName = sliceReg[2]

	return metricName, metricType

}

func New(storage storage.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {

	metricName, metricsType := UrlTreatmentGet(r.URL.Path)

	val, err := s.storage.Get(metricName, metricsType)
	if err == nil {
		fmt.Fprint(w, val)
		return
	}

	status := storageErrToStatus(err)
	w.WriteHeader(status)
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	var metricType, metricName, metricValue string

	metricType, metricName, metricValue = UrlTreatmentPush(r.URL.Path)

	err := s.storage.Put(metricName, metricValue, metricType)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	status := storageErrToStatus(err)
	w.WriteHeader(status)
}

func storageErrToStatus(err error) int {
	switch err {
	case storage.ErrAlreadyExists:
		return http.StatusConflict
	case storage.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
