package test

import (
	"awesomeDataRace/example_datarace/basic_datarace/example"
	"os"
	"testing"
)

// Случайно разделяемая переменная

func TestRaceAccidentalSharing(t *testing.T) {
	result := example.RaceAccidentalSharing(goroutines, iterations)
	if result != expected {
		t.Errorf("Possible data race due to accidental sharing. Expected %d, got %d", expected, result)
	} else {
		t.Logf("Expected %d, got %d", expected, result)
	}
}

func TestParallelWrite(t *testing.T) {
	data := []byte("Hello, race condition!")

	// Удалим файлы после теста
	defer func() {
		os.Remove("file1")
		os.Remove("file2")
	}()

	errs := example.ParallelWrite(data)

	// Проверим ошибки из двух горутин
	for i := 0; i < 2; i++ {
		err := <-errs
		if err != nil {
			t.Errorf("Error during write: %v", err)
		}
	}

	// Проверим, что оба файла созданы
	for _, fname := range []string{"file1", "file2"} {
		info, err := os.Stat(fname)
		if err != nil {
			t.Errorf("File %s not created: %v", fname, err)
			continue
		}
		if info.Size() != int64(len(data)) {
			t.Errorf("File %s has incorrect size: got %d, want %d", fname, info.Size(), len(data))
		}
	}
}
