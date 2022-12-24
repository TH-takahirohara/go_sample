package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := NewRouter()
	http.ListenAndServe(":8080", r)
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	r.Route("/test", func(r chi.Router) {
		r.Get("/{testId}", SampleServer)
	})
	return r
}

func SampleServer(w http.ResponseWriter, r *http.Request) {
	testId := chi.URLParam(r, "testId")
	w.Write([]byte(fmt.Sprintf("testId is %v", testId)))
}
