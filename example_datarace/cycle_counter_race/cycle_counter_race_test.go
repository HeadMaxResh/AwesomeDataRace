package cycle_counter_race

import (
	"testing"
)

// Гонка на счетчике цикла

const goroutines = 100
const iterations = 10000
const expected = goroutines * iterations

func TestRacePrimitive(t *testing.T) {
	result := RacePrimitive(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceObject(t *testing.T) {
	result := RaceObject(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceComposition(t *testing.T) {
	result := RaceComposition(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceAggregation(t *testing.T) {
	result := RaceAggregation(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestRaceSlice(t *testing.T) {
	result := RaceSlice(goroutines, iterations)
	if len(result) != expected {
		t.Errorf("Possible data race. Expected %d, got %d", expected, len(result))
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}
