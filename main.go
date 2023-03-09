package main

import (
	"context"
	"fmt"
	"github.com/dainoguchi/bstodo-api/internal/config"
	"github.com/dainoguchi/bstodo-api/internal/infra/auth0"
	"github.com/dainoguchi/bstodo-api/internal/infra/postgres"
	"github.com/dainoguchi/bstodo-api/internal/restapi/handler"
	"github.com/dainoguchi/bstodo-api/internal/restapi/middleware"
	"github.com/dainoguchi/bstodo-api/internal/usecase"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
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

	e := NewRouter(cfg, db)
	return e.Start(fmt.Sprintf(":%d", cfg.Port))
}

func NewRouter(cfg *config.Config, db *pgx.Conn) *echo.Echo {
	e := echo.New()

	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	// helth check
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	jv := auth0.NewJWTValidator(cfg.Auth0Domain, cfg.Auth0Audience)
	am := middleware.NewAuthMiddleware(jv)

	g := e.Group("/api/v1", am.EnsureValidToken)

	uh := handler.NewTodoHandler(usecase.NewTodoUsecase(db))
	g.POST("/todos", uh.Create)

	return e
}
