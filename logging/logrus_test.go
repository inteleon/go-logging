package logging_test

import (
	"encoding/json"
	"github.com/inteleon/go-logging/logging"
	"net"
	"testing"
)

type logFunc func(...interface{})

type jsonOutput struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
	Yo    string `json:"yo"`
	Mam   string `json:"mam"`
}

type testWriter struct {
	Value []byte
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.Value = p

	return
}

func initiateLogging(logLevel string) (logging.Logging, *testWriter) {
	w := &testWriter{}

	log, err := logging.NewLogrusLogging(logging.LogrusLoggingOptions{})

	if err != nil {
		panic(err)
	}

	log.SetOutput(w)
	log.SetLogLevel(logLevel)
	log.SetFormatter(logging.JSONFormatter)

	return log, w
}

func parseLogEntry(entry []byte) (out jsonOutput, err error) {
	err = json.Unmarshal(entry, &out)

	return
}

func assertPrinting(t *testing.T, logLevel string, f logFunc, writer *testWriter, extraVars bool) {
	expectedMessage := "hacker johnson"

	if extraVars {
		f(
			expectedMessage,
			map[string]string{
				"yo":  "lo",
				"mam": "ma",
			},
		)
	} else {
		f(expectedMessage, nil)
	}

	j, err := parseLogEntry(writer.Value)

	if err != nil {
		t.Fatal("Unexpected JSON decode error", err)
	}

	if j.Level != logLevel {
		t.Fatal("expected", logLevel, "got", j.Level)
	}

	if j.Msg != expectedMessage {
		t.Fatal("expected", expectedMessage, "got", j.Msg)
	}

	if extraVars {
		if j.Yo != "lo" {
			t.Fatal("expected", "lo", "got", j.Yo)
		}

		if j.Mam != "ma" {
			t.Fatal("expected", "ma", "got", j.Mam)
		}
	}
}

func TestDebugPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logging.DebugLogLevel)

	assertPrinting(t, logging.DebugLogLevel, log.Debug, writer, false)
}

func TestInfoPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logging.InfoLogLevel)

	assertPrinting(t, logging.InfoLogLevel, log.Info, writer, false)
}

func TestWarnPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logging.WarningLogLevel)

	assertPrinting(t, logging.WarningLogLevel, log.Warn, writer, false)
}

func TestErrorPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logging.ErrorLogLevel)

	assertPrinting(t, logging.ErrorLogLevel, log.Error, writer, false)
}

func TestDebugPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logging.DebugLogLevel)

	assertPrinting(t, logging.DebugLogLevel, log.Debug, writer, true)
}

func TestInfoPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logging.InfoLogLevel)

	assertPrinting(t, logging.InfoLogLevel, log.Info, writer, true)
}

func TestWarnPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logging.WarningLogLevel)

	assertPrinting(t, logging.WarningLogLevel, log.Warn, writer, true)
}

func TestErrorPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logging.ErrorLogLevel)

	assertPrinting(t, logging.ErrorLogLevel, log.Error, writer, true)
}

func TestSyslogNotRunningFailure(t *testing.T) {
	_, err := logging.NewLogrusLogging(
		logging.LogrusLoggingOptions{
			Syslog: &logging.LogrusLoggingSyslogOptions{
				Protocol: "tcp", // this test requires the tcp protocol
				Address:  "localhost:31337",
			},
		},
	)

	if err == nil {
		t.Fatal("expected", "error", "got", "nil")
	}
}

func TestSyslogRunningSuccess(t *testing.T) {
	go net.Listen("tcp", "localhost:31337")

	_, err := logging.NewLogrusLogging(
		logging.LogrusLoggingOptions{
			Syslog: &logging.LogrusLoggingSyslogOptions{
				Protocol: "tcp", // this test requires the tcp protocol
				Address:  "localhost:31337",
			},
		},
	)

	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}
