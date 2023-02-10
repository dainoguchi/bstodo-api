package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dainoguchi/bstodo-api/config"
	"github.com/dainoguchi/bstodo-api/infra/postgres"
	"github.com/dainoguchi/bstodo-api/restapi/handler"
	"github.com/dainoguchi/bstodo-api/usecase"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	cfg, _ := config.New()
	jst, err := time.LoadLocation(cfg.DBTZ)
	if err != nil {
		log.Fatal()
	}

	dsnConfig := postgres.DSNConfig{
		Host:     cfg.DBHost,
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		DbName:   cfg.DBName,
		Port:     cfg.DBPort,
		SslMode:  false,
		Loc:      jst,
	}

	db, err := postgres.New(dsnConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter(db)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}

func NewRouter(db *sql.DB) http.Handler {
	router := chi.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	// 試しにuser１件取得するのみ
	uh := handler.NewUserHandler(usecase.NewUserUsecase(db))
	router.HandleFunc("/user", uh.GetByID)

	return router
}
