package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dainoguchi/bstodo-api/internal/config"
	postgres "github.com/dainoguchi/bstodo-api/internal/infra/postgres"
	"github.com/dainoguchi/bstodo-api/internal/restapi/handler"
	"github.com/dainoguchi/bstodo-api/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
		DBName:   cfg.DBName,
		Port:     cfg.DBPort,
		SSLMode:  false,
		Loc:      jst,
	}

	db, err := postgres.New(dsnConfig)
	if err != nil {
		log.Fatal(err)
	}

	e := NewRouter(db)
	return e.Start(fmt.Sprintf(":%d", cfg.Port))
}

func NewRouter(db *sql.DB) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// 試しにuser１件取得するのみ
	uh := handler.NewUserHandler(usecase.NewUserUsecase(db))
	e.GET("/user", uh.GetByID)

	return e
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
