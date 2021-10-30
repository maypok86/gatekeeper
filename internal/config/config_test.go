package config

import (
	"os"
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
	}

	err := os.Setenv("LOGGER_LEVEL", testConfig.Logger.Level)
	require.NoError(t, err)
	err = os.Setenv("LOGGER_FILE", testConfig.Logger.File)
	require.NoError(t, err)

	cfg := Get()
	require.Equal(t, testConfig, cfg)
}
