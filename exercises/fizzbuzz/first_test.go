package main

import (
	"fmt"
	"testing"
)

func TestFizz(t *testing.T) {
	test_fizzbuzz := initCount(101)
	for i := range test_fizzbuzz {
		if i%3 == 0 {
			if test_fizzbuzz[i] == "Fizz" {
				t.Errorf("Fizz does not exist")
			} else {
				fmt.Printf("it exists")
			}
		}
	}
}
