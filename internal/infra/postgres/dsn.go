package postgres

import (
	"bytes"
	"fmt"
	"time"
)

const (
	defaultLocation       = "Asia/Tokyo"
	disabledSslModeString = "disable"
)

type DSNConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  bool
	Loc      *time.Location
}

// "host=db user=postgres password=password dbname=bstodo port=5432 sslmode=disable TimeZone=Asia/Tokyo")
func (cfg *DSNConfig) FormatDSN() string {
	// TODO: Validation

	var buf bytes.Buffer

	if len(cfg.Host) > 0 {
		buf.WriteString(fmt.Sprintf("host=%s ", cfg.Host))
	}

	if len(cfg.User) > 0 {
		buf.WriteString(fmt.Sprintf("user=%s ", cfg.User))
	}

	if len(cfg.Password) > 0 {
		buf.WriteString(fmt.Sprintf("password=%s ", cfg.Password))
	}

	if len(cfg.DBName) > 0 {
		buf.WriteString(fmt.Sprintf("dbname=%s ", cfg.DBName))
	}

	if cfg.Port != 0 {
		buf.WriteString(fmt.Sprintf("port=%d ", cfg.Port))
	}

	if !cfg.SSLMode {
		buf.WriteString(fmt.Sprintf("sslmode=%s ", disabledSslModeString))
	}

	// defaultLocation Asia/Tokyoに設定しちゃってるけどUTC？普通どこだろ
	if len(cfg.Loc.String()) > 0 {
		buf.WriteString(fmt.Sprintf("TimeZone=%s", cfg.Loc.String()))
	} else {
		buf.WriteString(fmt.Sprintf("TimeZone=%s", defaultLocation))
	}

	return buf.String()
}
