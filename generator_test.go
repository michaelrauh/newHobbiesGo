package main

import "testing"

func TestThatGenerateReturnsAStringOfSpecifiedLength(t *testing.T) {
	actual := Generate(6)
	if len(actual) != 6 {
		t.Errorf("Generate created a string of the wrong length, got: %s, want: %d.", actual, 6)
	}
}
