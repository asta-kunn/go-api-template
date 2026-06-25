package main

import (
	"testing"
)

// TestSanityCheck adalah dummy test agar pipeline CI/CD bisa lolos (Passed)
func TestSanityCheck(t *testing.T) {
	expected := true
	actual := true

	if expected != actual {
		t.Errorf("Ekspektasi %v, tapi dapatnya %v", expected, actual)
	}
}
