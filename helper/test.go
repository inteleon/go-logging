package helper

import (
	"encoding/json"
	"github.com/inteleon/go-logging/logging"
	"testing"
)

// TestLoggingJSONOutput is the expected JSON logging output structure.
type TestLoggingJSONOutput struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

// TestLogWriter is a test log writer that only collects what is written to it, doesn't output anything.
type TestLogWriter struct {
	Buffer [][]byte
}

// Write appends what is written to the in-memory buffer.
func (s *TestLogWriter) Write(p []byte) (count int, err error) {
	s.Buffer = append(s.Buffer, p)

	return
}

// NewTestLogging creates and returns a new test logging object.
func NewTestLogging() (logging.Logging, *TestLogWriter) {
	w := &TestLogWriter{
		Buffer: [][]byte{},
	}

	l, err := logging.NewLogrusLogging(logging.LogrusLoggingOptions{})

	if err != nil {
		panic(err) // For backwards compatability - so we don't need to return the err value and handle it that way.
	}

	l.SetOutput(w)
	l.SetLogLevel(logging.DebugLogLevel)
	l.SetFormatter(logging.JSONFormatter)

	return l, w
}

// ParseLogEntry takes care of parsing a log entry into a JSON structure.
func ParseLogEntry(entry []byte) (out TestLoggingJSONOutput, err error) {
	err = json.Unmarshal(entry, &out)

	return
}

// ValidateLogEntry is a helper tests can use to validate simple logging entries.
func ValidateLogEntry(t *testing.T, entry []byte, logLevel string, expectedMsg string) {
	j, err := ParseLogEntry(entry)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if j.Level != logLevel {
		t.Fatal("expected", logLevel, "got", j.Level)
	}

	if j.Msg != expectedMsg {
		t.Fatal("expected", expectedMsg, "got", j.Msg)
	}
}
