package logger

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_Logger(t *testing.T) {
	t.Run("logging", func(t *testing.T) {
		filename := "develop.log"
		Configure(&config.Logger{
			File:  filename,
			Level: "info",
		})
		zap.L().Info("logger construction succeeded")

		file, err := ioutil.ReadFile(filename)
		require.NoError(t, err)
		got := string(file)
		err = os.Remove(filename)
		require.NoError(t, err)

		require.Contains(t, got, "logger construction succeeded\n")
	})
}
