package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func New(dsnConfig DSNConfig) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsnConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to create db conn: %s", err.Error())
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	return conn, nil
}
