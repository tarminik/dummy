package compose

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// captureOutput redirects os.Stdout to a buffer during function f, then returns the output.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = old
	return buf.String()
}

func TestUp(t *testing.T) {
	output := captureOutput(func() { Up("svc", "path/to/config.yaml") })
	if !strings.Contains(output, "Starting environment for service 'svc'") {
		t.Error("Up() did not print start message")
	}
}

func TestDown(t *testing.T) {
	output := captureOutput(func() { Down("svc", "path/to/config.yaml") })
	if !strings.Contains(output, "Stopping environment for service 'svc'") {
		t.Error("Down() did not print stop message")
	}
}

func TestLogs(t *testing.T) {
	output := captureOutput(func() { Logs("svc") })
	if !strings.Contains(output, "Logs for service 'svc'") {
		t.Error("Logs() did not print logs header")
	}
	if !strings.Contains(output, "[INFO] Service started successfully.") {
		t.Error("Logs() did not print info log")
	}
}

func TestStatus(t *testing.T) {
	output := captureOutput(func() { Status("svc") })
	if !strings.Contains(output, "Service svc: Running (simulation)") {
		t.Error("Status() did not print running status")
	}
}

func TestRunTests(t *testing.T) {
	output := captureOutput(func() { RunTests("svc", "pytest") })
	if !strings.Contains(output, "Running tests for service 'svc' with command: pytest") {
		t.Error("RunTests() did not print running tests message")
	}
	if !strings.Contains(output, "Tests finished: PASSED (simulation)") {
		t.Error("RunTests() did not print test finished message")
	}
}
