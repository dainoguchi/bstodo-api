package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
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

type PgxWrapper interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

var (
	_ PgxWrapper = (*pgx.Conn)(nil)
	_ PgxWrapper = (pgx.Tx)(nil)
)
