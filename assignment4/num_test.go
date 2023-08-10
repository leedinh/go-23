package main

import "testing"

func TestNumDifferentInteger(t *testing.T) {
	word := "A1b01c001"

	count := numDifferentInteger(word)
	if count != 1 {
		t.Error("Expected 1, got ", count)
	}

	word = "leet1234code234"

	count = numDifferentInteger(word)
	if count != 2 {
		t.Error("Expected 2, got ", count)
	}
}
