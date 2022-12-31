package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to run server: %v", err)
	}
}

func run(ctx context.Context) error {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")),
	)
	if err != nil {
		return fmt.Errorf("failed open db: %w", err)
	}
	defer db.Close()

	r := NewRouter(ctx, db)
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

func NewRouter(ctx context.Context, db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("health ok"))
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		var id int
		var name string
		rows, err := db.Query("select id, name from users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	})

	return r
}
