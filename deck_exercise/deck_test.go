package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected lenght of 52 but got %d", len(cards))
	}

	if cards[0] != "1 of Spades" {
		t.Errorf("Expected the 1 of Spades as the outcome but got %v", cards[0])
	}

	if cards[len(cards)-1] != "K of Clubs" {
		t.Errorf("Expected the K of Clubs as the outcome but got %v", cards[len(cards)-1])
	}
}

func TestSaveReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	newDeck := readFromFile("_decktesting")

	if len(newDeck) != 52 {
		t.Errorf("Expected lenght of 52 but got %d", len(newDeck))
	}

	os.Remove("_decktesting")
}
