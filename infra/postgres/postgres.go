package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// TODO: 環境変数に詰め替える
func New() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=db user=postgres password=password dbname=bstodo port=5432 sslmode=disable TimeZone=Asia/Tokyo")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create db conn: %s", err.Error())
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	return conn, nil
}
