package test

import (
	"awesomeDataRace/example_datarace/basic_datarace/example"
	"testing"
)

// Гонка на счетчике цикла

func TestRacePrimitive(t *testing.T) {
	result := example.RacePrimitive(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceObject(t *testing.T) {
	result := example.RaceObject(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceComposition(t *testing.T) {
	result := example.RaceComposition(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceAggregation(t *testing.T) {
	result := example.RaceAggregation(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceSlice(t *testing.T) {
	result := example.RaceSlice(goroutines, iterations)
	if len(result) != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, len(result))
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}
