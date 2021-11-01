package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/maypok86/gatekeeper/pkg/logger"
)

type Config struct {
	Logger *logger.Config
	Bucket *Bucket
}

type Bucket struct {
	Type          string        `envconfig:"BUCKET_TYPE" default:"inmemory"`
	TTL           time.Duration `envconfig:"BUCKET_TTL" default:"60ns"`
	CleanInterval time.Duration `envconfig:"BUCKET_CLEAN_INTERVAL" default:"61ns"`
	Rate          int           `envconfig:"BUCKET_RATE"`
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
