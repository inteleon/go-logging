package logging_test

import (
	"encoding/json"
	"testing"

	"github.com/inteleon/go-logging/logging"
)

func testLogger(level, activeLevel string, hasOutput bool, t *testing.T) {
	w := &testWriter{}
	log, err := logging.NewLogrusLogging(logging.LogrusLoggingOptions{})

	if err != nil {
		t.Fatal(err)
	}

	log.SetOutput(w)
	log.SetLogLevel(activeLevel)
	log.SetFormatter(logging.JSONFormatter)

	logger := log.Logger(level)

	logger.Printf("zomg%s", "lol")

	if hasOutput {
		var output jsonOutput
		err = json.Unmarshal(w.Value, &output)

		if err != nil {
			t.Error(err)
		}

		if output.Msg != "zomglol" {
			t.Error("Expected zomglol, got ", output.Msg)
		}
	} else {
		if w.Value != nil {
			t.Error("Expected nil")
		}
	}
}

func TestDebugLogger(t *testing.T) {
	testLogger(logging.DebugLogLevel, logging.DebugLogLevel, true, t)
	testLogger(logging.DebugLogLevel, logging.InfoLogLevel, false, t)
}

func TestInfoLogger(t *testing.T) {
	testLogger(logging.InfoLogLevel, logging.InfoLogLevel, true, t)
	testLogger(logging.InfoLogLevel, logging.WarningLogLevel, false, t)
}

func TestWarningLogger(t *testing.T) {
	testLogger(logging.WarningLogLevel, logging.WarningLogLevel, true, t)
	testLogger(logging.InfoLogLevel, logging.WarningLogLevel, false, t)
}

func TestErrorLogger(t *testing.T) {
	testLogger(logging.ErrorLogLevel, logging.ErrorLogLevel, true, t)
	testLogger(logging.ErrorLogLevel, logging.FatalLogLevel, false, t)
}
