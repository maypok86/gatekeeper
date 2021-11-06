package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBType       string `envconfig:"DB_TYPE" default:"inmemory"`
	Host         string `envconfig:"HOST" default:"0.0.0.0"`
	Port         string `envconfig:"PORT" default:"50051"`
	RateLogin    int    `envconfig:"RATE_LOGIN" default:"10"`
	RatePassword int    `envconfig:"RATE_PASSWORD" default:"100"`
	RateIP       int    `envconfig:"RATE_IP" default:"1000"`
	Logger       *Logger
}

type Logger struct {
	Level string `envconfig:"LOGGER_LEVEL" default:"info"`
	File  string `envconfig:"LOGGER_FILE" default:"develop.log"`
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

func GetBucketByRate(rate int) *Bucket {
	return &Bucket{
		Rate:          rate,
		TTL:           defaultTTL,
		CleanInterval: defaultCleanInterval,
	}
}

func Get() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
