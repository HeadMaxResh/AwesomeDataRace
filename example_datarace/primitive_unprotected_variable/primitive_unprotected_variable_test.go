package primitive_unprotected_variable

import "testing"

// Примитивная незащищенная переменная

const goroutines = 100
const iterations = 10000
const expected = goroutines * iterations

func TestRacePrimitiveBad(t *testing.T) {
	result := RacePrimitive(goroutines, iterations)
	if result != expected {
		t.Errorf("Data race likely occurred. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}
