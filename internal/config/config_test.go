package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/maypok86/gatekeeper/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestConfig_Get(t *testing.T) {
	testConfig := &Config{
		Logger: &logger.Config{
			Level: "warn",
			File:  "test.log",
		},
		Bucket: &Bucket{
			TTL:           61,
			CleanInterval: 62,
			Rate:          20,
		},
	}

	err := os.Setenv("LOGGER_LEVEL", testConfig.Logger.Level)
	require.NoError(t, err)
	err = os.Setenv("LOGGER_FILE", testConfig.Logger.File)
	require.NoError(t, err)
	err = os.Setenv("BUCKET_TYPE", testConfig.Bucket.Type)
	require.NoError(t, err)
	err = os.Setenv("BUCKET_TTL", testConfig.Bucket.TTL.String())
	require.NoError(t, err)
	err = os.Setenv("BUCKET_CLEAN_INTERVAL", testConfig.Bucket.CleanInterval.String())
	require.NoError(t, err)
	err = os.Setenv("BUCKET_RATE", strconv.Itoa(testConfig.Bucket.Rate))
	require.NoError(t, err)

	cfg := Get()
	require.Equal(t, testConfig, cfg)
}
