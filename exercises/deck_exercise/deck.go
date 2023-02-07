package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of strings
type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// function to create/initialize a new Deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}

// function to deal a deck based on the handsize given from the user
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save the deck to a file locally with permissions 0666
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// read from a file stored locally and pass the content to a new deck
func readFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error while reading the file. Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle the deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
