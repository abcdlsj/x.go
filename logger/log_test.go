package logger

import (
	"os"
	"testing"
)

func defaultTestPretty(f *os.File) *PrettyHandler {
	return NewPrettyHandler().W(f).SetLevel(Info)
}

func TestPrintInfo(t *testing.T) {
	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var xlogger = New(WithHandler(defaultTestPretty(f)))
	xlogger.Info("run logger", "name", "logger", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}

func TestPrintDebug(t *testing.T) {
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var xlogger = New(WithHandler(defaultTestPretty(f)))
	xlogger.Debug("run logger", "name", "logger", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}

func TestPrintError(t *testing.T) {
	f, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var xlogger = New(WithHandler(defaultTestPretty(f)))
	xlogger.Error("run logger", "name", "logger", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}

func TestPrintWarn(t *testing.T) {
	f, err := os.OpenFile("warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var xlogger = New(WithHandler(defaultTestPretty(f)))
	xlogger.Warn("run logger", "name", "logger", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}

func TestMarkHandler(t *testing.T) {
	f, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	markFields := map[string]interface{}{
		"version": "v*.*.*",
	}
	var xlogger = New(WithHandler(NewMarkHandler(markFields), defaultTestPretty(f)))
	xlogger.Info("run logger", "name", "logger", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}
