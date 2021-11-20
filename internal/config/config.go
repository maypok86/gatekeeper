package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvType string

const (
	test EnvType = "test"
	prod EnvType = "prod"
	dev  EnvType = "dev"
)

type Config struct {
	Environment    EnvType `envconfig:"ENVIRONMENT"`
	DBType         string  `envconfig:"DB_TYPE"`
	Host           string  `envconfig:"HOST"`
	Port           string  `envconfig:"PORT"`
	RateLogin      int     `envconfig:"RATE_LOGIN"`
	RatePassword   int     `envconfig:"RATE_PASSWORD"`
	RateIP         int     `envconfig:"RATE_IP"`
	PrometheusPort string  `envconfig:"PROMETHEUS_PORT"`
	Logger         *Logger
}

type Logger struct {
	Level string `envconfig:"LOGGER_LEVEL"`
	File  string `envconfig:"LOGGER_FILE"`
}

type Bucket struct {
	TTL           time.Duration
	CleanInterval time.Duration
	Rate          int
}

const (
	defaultTTL           = 60
	defaultCleanInterval = 61
)

var (
	config Config
	once   sync.Once
)

func (c *Config) IsDev() bool {
	return c.Environment == dev
}

func GetBucketByRate(rate int) *Bucket {
	return &Bucket{
		Rate:          rate,
		TTL:           defaultTTL,
		CleanInterval: defaultCleanInterval,
	}
}

func Get() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Print("Error loading .env file")
		}
		if err := envconfig.Process("", &config); err != nil {
			log.Fatal(err)
		}
		switch config.Environment {
		case test, prod, dev:
		default:
			log.Fatal("config environment should be test, prod or dev")
		}
		configBytes, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
