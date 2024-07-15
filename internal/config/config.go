package config

import (
	"errors"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"time"
)

type AppConfig struct {
	Name string `env:"APP_NAME" envDefault:"time_tracker_db"`
}

type HTTPConfig struct {
	Host         string        `env:"HOST" envDefault:"localhost"`
	Port         string        `env:"PORT" envDefault:"3000"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
}

type ThirdPartyService struct {
	Url string `env:"API_URL"`
}

type SQLConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	DBName   string `env:"DB_NAME" envDefault:"time_tracker_db"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SslMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
}

type Config struct {
	App     AppConfig
	HTTP    HTTPConfig
	Db      SQLConfig
	UserApi ThirdPartyService
}

func NewConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Error("error loading env variables: %s", err.Error())
		return Config{}, errors.New("error loading env variables")
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("parse error config ", err.Error())
		return cfg, err
	}
	return cfg, nil
}
