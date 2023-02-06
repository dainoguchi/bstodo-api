package config

import "github.com/caarlos0/env/v7"

type config struct {
	Port   int    `env:"PORT" envDefault:"8080"`
	DBName string `env:"DB_NAME" envDefault:"bstodo"`
	DBPort int    `env:"DB_PORT" envDefault:"5432"`
	DBPass string `env:"DB_PASS" envDefault:"password"`
	DBUser string `env:"DB_USER" envDefault:"postgres"`
	DBHost string `env:"DB_HOST" envDefafult:"127.0.0.1"`
}

func New() (*config, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
