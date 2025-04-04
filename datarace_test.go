package main

import (
	"testing"
)

func TestRacePrimitive(t *testing.T) {
	result := RacePrimitive()
	if result != 2000 {
		t.Errorf("Possible data race. Expected 2000, got %d", result)
	} else {
		t.Logf("Expected 2000, got %d", result)
	}
}

func TestRaceObject(t *testing.T) {
	result := RaceObject()
	if result != 2000 {
		t.Errorf("Possible data race. Expected 2000, got %d", result)
	} else {
		t.Logf("Expected 2000, got %d", result)
	}
}

func TestRaceComposition(t *testing.T) {
	result := RaceComposition()
	if result != 2000 {
		t.Errorf("Possible data race. Expected 2000, got %d", result)
	} else {
		t.Logf("Expected 2000, got %d", result)
	}
}

func TestRaceAggregation(t *testing.T) {
	result := RaceAggregation()
	if result != 2000 {
		t.Errorf("Possible data race. Expected 2000, got %d", result)
	} else {
		t.Logf("Expected 2000, got %d", result)
	}
}

func TestRaceSlice(t *testing.T) {
	result := RaceSlice()
	if len(result) != 2000 {
		t.Errorf("Possible data race. Expected 2000 elements, got %d", len(result))
	} else {
		t.Logf("Expected 2000, got %d", result)
	}
}
