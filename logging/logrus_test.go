package logging_test

import (
	"encoding/json"
	"github.com/inteleon/go-logging/logging"
	logrus "github.com/sirupsen/logrus"
	"testing"
)

type logFunc func(string, interface{})

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

func initiateLogging(logLevel logrus.Level) (logging.Logging, *testWriter) {
	w := &testWriter{}

	log := logging.NewLogrusLogging(
		logLevel,
		w,
		&logrus.JSONFormatter{},
	)

	return log, w
}

func parseLogEntry(entry []byte) (out jsonOutput, err error) {
	err = json.Unmarshal(entry, &out)

	return
}

func assertPrinting(t *testing.T, logLevel logrus.Level, f logFunc, writer *testWriter, extraVars bool) {
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

	if j.Level != logLevel.String() {
		t.Fatal("expected", logLevel.String(), "got", j.Level)
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
	log, writer := initiateLogging(logrus.DebugLevel)

	assertPrinting(t, logrus.DebugLevel, log.Debug, writer, false)
}

func TestInfoPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.InfoLevel)

	assertPrinting(t, logrus.InfoLevel, log.Info, writer, false)
}

func TestWarnPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.WarnLevel)

	assertPrinting(t, logrus.WarnLevel, log.Warn, writer, false)
}

func TestErrorPrintingWithoutVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.ErrorLevel)

	assertPrinting(t, logrus.ErrorLevel, log.Error, writer, false)
}

func TestDebugPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.DebugLevel)

	assertPrinting(t, logrus.DebugLevel, log.Debug, writer, true)
}

func TestInfoPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.InfoLevel)

	assertPrinting(t, logrus.InfoLevel, log.Info, writer, true)
}

func TestWarnPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.WarnLevel)

	assertPrinting(t, logrus.WarnLevel, log.Warn, writer, true)
}

func TestErrorPrintingWithVariables(t *testing.T) {
	log, writer := initiateLogging(logrus.ErrorLevel)

	assertPrinting(t, logrus.ErrorLevel, log.Error, writer, true)
}
