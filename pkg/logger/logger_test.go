package logger

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Logger(t *testing.T) {
	t.Run("logging", func(t *testing.T) {
		filename := "develop.log"
		Configure(&Config{
			File:  filename,
			Level: "info",
		})
		Info("logger construction succeeded")

		file, err := ioutil.ReadFile(filename)
		require.NoError(t, err)
		got := string(file)
		err = os.Remove(filename)
		require.NoError(t, err)

		require.Contains(t, got, "logger construction succeeded\n")
	})
}
