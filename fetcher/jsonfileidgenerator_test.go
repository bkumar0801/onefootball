package fetcher

import (
	"testing"
)

func TestIdGeneratorOperation(t *testing.T) {
	generator := NewIDGenerator(1, 6, 10)

	cases := []struct {
		currentID int
		isMaxID   bool
	}{
		{1, true},
		{7, false},
	}

	for _, c := range cases {
		actualID := generator.Current()
		if actualID != c.currentID {
			t.Errorf("Current expected: %q, actual: %q", c.currentID, actualID)
		}
		actualStatus := generator.GenerateNext()
		if actualStatus != c.isMaxID {
			t.Errorf("Current expected: %t, actual: %t", c.isMaxID, actualStatus)
		}
	}
}
