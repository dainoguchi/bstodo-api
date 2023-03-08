package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func New(dsnConfig DSNConfig) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsnConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to create db conn: %s", err.Error())
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	return conn, nil
}
