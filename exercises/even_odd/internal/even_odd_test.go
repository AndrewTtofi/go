package internal

import "testing"

func TestEvenOdd(t *testing.T) {
	output := SpecificNum(5)
	expected := "The number 5 is odd"
	if output != expected {
		t.Errorf("got %s but wanted %s", output, expected)
	}
}
