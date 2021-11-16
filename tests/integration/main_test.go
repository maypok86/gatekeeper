package integration

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                 "integration",
		TestSuiteInitializer: nil,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:    "progress",
			Paths:     []string{"features"},
			Randomize: 0,
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
