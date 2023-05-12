package xlog

import (
	"os"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var xlogger = New(WithHandler(NewPrettyHandler(f, true)))
	xlogger.Info("run xlog", "name", "xlog", "version", "v1.0.0")
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
	var xlogger = New(WithHandler(NewPrettyHandler(f, true)))
	xlogger.Debug("run xlog", "name", "xlog", "version", "v1.0.0")
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
	var xlogger = New(WithHandler(NewPrettyHandler(f, true)))
	xlogger.Error("run xlog", "name", "xlog", "version", "v1.0.0")
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
	var xlogger = New(WithHandler(NewPrettyHandler(f, true)))
	xlogger.Warn("run xlog", "name", "xlog", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}

func TestMarkHandler(t *testing.T) {
	f, err := os.OpenFile("warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var xlogger = New(WithHandler(NewMarkHandler(map[string]interface{}{"version": "****"}), NewPrettyHandler(f, true)))
	xlogger.Info("run xlog", "name", "xlog", "version", "v1.0.0")
	fdata, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("log file content: %s", fdata)
}
