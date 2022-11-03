package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string //create a new type of deck that will be basically a slice of type string

func newDeck() deck {
	cards := deck{} //Initialize the cards slice of type deck
	cardtypes := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardnumbers := []string{"Ace", "Two", "Three", "Four"}
	for _, cardtype := range cardtypes {
		for _, cardnumber := range cardnumbers {
			cards = append(cards, cardnumber+" of "+cardtype)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //transform the slice into a single line string and set the delimeter "," that wil seperate the value of the slice
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //bs stands for byteslice and will be the value of the file
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}
