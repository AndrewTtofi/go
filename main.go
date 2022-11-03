package main

func main() {
	// cards := newDeck()
	// // cards.print()
	// hand, remainingDeck := deal(cards, 4)
	// hand.print()
	// remainingDeck.print()
	// cards.print()
	// cards.saveToFile("giannakis")
	cards := newDeckFromFile("giannakis")
	cards.print()
}
