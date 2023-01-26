package server

import (
	"fmt"
	"net/http"

	"fraima.io/fraimmon/internal/storage"
	"fraima.io/fraimmon/internal/util"
)

type Server struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	m, code := util.UrlTreatment(r.URL.Path)
	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	val, code := s.storage.Get(m)
	if code == http.StatusOK {
		fmt.Fprint(w, val)
		return
	}

	w.WriteHeader(code)

}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	m, code := util.UrlTreatment(r.URL.Path)
	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	code = s.storage.Put(m)
	if code == http.StatusOK {
		w.WriteHeader(code)
		return
	}

	w.WriteHeader(code)

}
