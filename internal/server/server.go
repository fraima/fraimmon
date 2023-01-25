package server

import (
	"fmt"
	"net/http"
	"regexp"

	"fraima.io/fraimmon/internal/storage"
	"fraima.io/fraimmon/internal/types"
)

type Server struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func UrlTreatment(uri string) (types.MetricItem, int) {

	// log.Printf("<UrlTreatment> start func")
	// log.Printf("<UrlTreatment> init payload: %s", uri)

	var m types.MetricItem
	re := regexp.MustCompile(`\/update\/(counter|gauge)\/(\w*)\/(\w.*)`)
	sliceReg := re.FindStringSubmatch(uri)

	if len(sliceReg) == 4 {
		m.Type = sliceReg[1]
		m.Name = sliceReg[2]
		m.Value = sliceReg[3]

	} else if len(sliceReg) == 3 {
		m.Type = sliceReg[1]
		m.Name = sliceReg[2]
		m.Value = ""

	} else {
		return m, http.StatusNotFound
	}

	return m, http.StatusOK

}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	// log.Printf("<Get:server> start func")
	var m types.MetricItem
	m, code := UrlTreatment(r.URL.Path)

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	val, code := s.storage.Get(m)
	if code == http.StatusOK {
		fmt.Fprint(w, val)
		return
	}

	// status := problem.StorageErrToStatus(err)
	w.WriteHeader(code)
	// log.Printf("<Get:server> exit func")
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	// log.Printf("<Put:server> start func")
	var m types.MetricItem

	m, _ = UrlTreatment(r.URL.Path)

	// log.Printf("<Put:server> payload <- <UrlTreatment>: %s", m)
	code := s.storage.Put(m)

	if code == http.StatusOK {
		w.WriteHeader(code)
		return
	} //else {
	// log.Printf("<Put:server> status code: %s", strconv.FormatInt(int64(code), 10))
	//}

	w.WriteHeader(code)
	// log.Printf("<Put:server> exit func")
}
