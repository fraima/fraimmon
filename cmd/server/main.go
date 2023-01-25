package main

import (
	"net/http"

	"fraima.io/fraimmon/internal/server"
	"fraima.io/fraimmon/internal/storage"
	"github.com/go-chi/chi"
)

func main() {
	// маршрутизация запросов обработчику

	m := chi.NewRouter()

	st := storage.NewInMemory()

	s := server.New(st)

	m.Get("/update/*", s.Get)
	m.Post("/update/*", s.Put)

	http.ListenAndServe(":8080", m)
}
