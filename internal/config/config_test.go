package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_Get(t *testing.T) {
	testConfig := &Config{
		DBType: "inmemory",
		Logger: &Logger{
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
	err = os.Setenv("DB_TYPE", testConfig.DBType)
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
