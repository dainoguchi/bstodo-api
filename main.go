package main

import (
	"context"
	"fmt"
	"github.com/dainoguchi/bstodo-api/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	cfg, _ := config.New()
	router := NewRouter()

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	return router
}
