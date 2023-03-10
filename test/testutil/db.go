package testutil

import (
	"github.com/dainoguchi/bstodo-api/internal/config"
	"github.com/dainoguchi/bstodo-api/internal/infra/postgres"
	"github.com/jackc/pgx/v4"
	"testing"
	"time"
)

func OpenDBForTest(t *testing.T) *pgx.Conn {
	t.Helper()

	cfg, _ := config.New()
	jst, err := time.LoadLocation(cfg.DBTZ)
	if err != nil {
		t.Fatal(err)
	}

	dsnConfig := postgres.DSNConfig{
		Host:     cfg.DBHost,
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		DBName:   cfg.DBNameTest,
		Port:     cfg.DBPort,
		SSLMode:  false,
		Loc:      jst,
	}

	db, err := postgres.New(dsnConfig)
	if err != nil {
		t.Fatalf("create db error %v\n dsn: %s", err, dsnConfig.FormatDSN())
	}

	// TODO: priorityに初期値が挿入されているか

	return db
}
