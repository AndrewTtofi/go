package main

import (
	"fmt"
	"testing"
)

func TestFizz(t *testing.T) {
	testFizzbuzz := initCount(101)
	for i := range testFizzbuzz {
		if i%3 == 0 {
			if testFizzbuzz[i] == "Fizz" {
				t.Errorf("Fizz does not exist")
			} else {
				fmt.Printf("it exists")
			}
		}
	}
}
