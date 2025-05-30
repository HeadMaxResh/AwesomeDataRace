package test

import (
	"awesomeDataRace/example_datarace/basic_datarace/example"
	"testing"
)

// Незащищенная глобальная переменная

func TestRaceGlobal(t *testing.T) {
	result := example.RaceGlobal(goroutines, iterations)
	if result != expected {
		t.Errorf("Data race likely occurred. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}
