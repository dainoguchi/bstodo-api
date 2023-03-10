package config

import "github.com/caarlos0/env/v7"

type Config struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	DBName     string `env:"DB_NAME" envDefault:"bstodo"`
	DBNameTest string `env:"DB_NAME_TEST" envDefault:"bstodo_test"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBPass     string `env:"DB_PASS" envDefault:"password"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBTZ       string `env:"DB_TZ"   envDefault:"Asia/Tokyo"`

	Auth0Domain   string `env:"AUTH0_DOMAIN"`
	Auth0Audience string `env:"AUTH0_AUDIENCE"`
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
