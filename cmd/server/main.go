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

	m.Get("/update/counter/*", s.Get)
	m.Post("/update/counter/*", s.Put)

	m.Get("/update/gauge/*", s.Get)
	m.Post("/update/gauge/*", s.Put)

	http.ListenAndServe(":8080", m)
}
