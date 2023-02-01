package main

import "testing"

func TestMain(t *testing.T) {
	total := Multiply()
	if total != 100 {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 100)
	}
}
