package randomly_shared_variable

import (
	"testing"
)

// Случайно разделяемая переменная

const goroutines = 100
const iterations = 10000
const expected = goroutines * iterations

func TestRaceAccidentalSharing(t *testing.T) {
	result := RaceAccidentalSharing(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race due to accidental sharing. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}
