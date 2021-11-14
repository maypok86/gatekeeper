package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_Get(t *testing.T) {
	testConfig := &Config{
		Environment:    "prod",
		DBType:         "inmemory",
		Host:           "0.0.0.1",
		Port:           "50050",
		RateLogin:      11,
		RatePassword:   101,
		RateIP:         1001,
		PrometheusPort: "9090",
		Logger: &Logger{
			Level: "warn",
			File:  "test.log",
		},
	}

	err := os.Setenv("LOGGER_LEVEL", testConfig.Logger.Level)
	require.NoError(t, err)
	err = os.Setenv("LOGGER_FILE", testConfig.Logger.File)
	require.NoError(t, err)
	err = os.Setenv("ENVIRONMENT", string(testConfig.Environment))
	require.NoError(t, err)
	err = os.Setenv("DB_TYPE", testConfig.DBType)
	require.NoError(t, err)
	err = os.Setenv("HOST", testConfig.Host)
	require.NoError(t, err)
	err = os.Setenv("PORT", testConfig.Port)
	require.NoError(t, err)
	err = os.Setenv("RATE_LOGIN", strconv.Itoa(testConfig.RateLogin))
	require.NoError(t, err)
	err = os.Setenv("RATE_PASSWORD", strconv.Itoa(testConfig.RatePassword))
	require.NoError(t, err)
	err = os.Setenv("RATE_IP", strconv.Itoa(testConfig.RateIP))
	require.NoError(t, err)
	err = os.Setenv("PROMETHEUS_PORT", testConfig.PrometheusPort)
	require.NoError(t, err)

	cfg := Get()
	require.Equal(t, testConfig, cfg)
}
