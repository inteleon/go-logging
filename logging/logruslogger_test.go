package logging

import (
	"encoding/json"
	"testing"
)

type testWriter struct {
	Value []byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.Value = p
	return
}

type jsonOutput struct {
	Msg string `json:"msg"`
}

func testLogger(level, activeLevel string, hasOutput bool, t *testing.T) {
	w := &testWriter{}
	log, err := NewLogrusLogging(LogrusLoggingOptions{})

	if err != nil {
		panic(err)
	}

	log.SetOutput(w)
	log.SetLogLevel(activeLevel)
	log.SetFormatter(JSONFormatter)

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
	testLogger(DebugLogLevel, DebugLogLevel, true, t)
	testLogger(DebugLogLevel, InfoLogLevel, false, t)
}

func TestInfoLogger(t *testing.T) {
	testLogger(InfoLogLevel, InfoLogLevel, true, t)
	testLogger(InfoLogLevel, WarningLogLevel, false, t)
}

func TestWarningLogger(t *testing.T) {
	testLogger(WarningLogLevel, WarningLogLevel, true, t)
	testLogger(WarningLogLevel, ErrorLogLevel, false, t)
}

func TestErrorLogger(t *testing.T) {
	testLogger(ErrorLogLevel, ErrorLogLevel, true, t)
	testLogger(ErrorLogLevel, FatalLogLevel, false, t)
}
