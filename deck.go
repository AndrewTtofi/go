package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
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
