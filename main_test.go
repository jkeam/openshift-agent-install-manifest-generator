package main

import (
	"testing"
)

func TestMath(t *testing.T) {
	a := 1 + 3
	if a != 4 {
		t.Errorf("Math is broken")
	}
}
