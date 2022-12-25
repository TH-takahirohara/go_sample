package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to run server: %v", err)
	}
}

func run(ctx context.Context) error {
	r := NewRouter(ctx)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	err = http.Serve(l, r)
	if err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func NewRouter(ctx context.Context) *chi.Mux {
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
