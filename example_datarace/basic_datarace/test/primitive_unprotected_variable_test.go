package test

import (
	"awesomeDataRace/example_datarace/basic_datarace/example"
	"testing"
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
