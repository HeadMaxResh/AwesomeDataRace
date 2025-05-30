package test

import (
	"awesomeDataRace/example_datarace/basic_datarace/example"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// Примитивная незащищенная переменная

func TestRaceUnprotectedPrimitive(t *testing.T) {
	result := example.RacePrimitive(goroutines, iterations)
	if result != expected {
		t.Errorf("Data race likely occurred. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestWatchdogRace_DetectsRace(t *testing.T) {
	cmd := exec.Command("go", "test", "-race", "-run=TestWatchdogInner")
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()

	if err == nil {
		t.Fatal("Expected race detector to fail, but test passed.")
	}

	if !strings.Contains(string(out), "WARNING: DATA RACE") {
		t.Fatalf("Expected race warning, got:\n%s", out)
	}
}

func TestWatchdogInner(t *testing.T) {
	w := &example.Watchdog{}
	w.Start()
	for i := 0; i < 1000; i++ {
		w.KeepAlive()
		time.Sleep(1 * time.Millisecond)
	}
}
