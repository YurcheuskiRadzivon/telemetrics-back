package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App   App
		HTTP  HTTP
		PG    PG
		REDIS REDIS
		TG    TG
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	// DB -.
	PG struct {
		URL string `env:"PG_URL,required"`
	}

	// REDIS
	REDIS struct {
		URL string `env:"REDIS_URL,required"`
	}

	//TG
	TG struct {
		APP_ID   int    `env:"APP_ID,required"`
		APP_HASH string `env:"APP_HASH,required"`
	}
)

// NewConfig returns app config
func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
