package unprotected_global_variable

import "testing"

// Незащищенная глобальная переменная

const goroutines = 100
const iterations = 10000
const expected = goroutines * iterations

func TestRaceGlobalBad(t *testing.T) {
	result := RaceGlobal(goroutines, iterations)
	if result != expected {
		t.Errorf("Data race likely occurred. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}
