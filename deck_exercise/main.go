package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("this is the hand")
	// hand.print()
	// fmt.Println("this is the remaining deck")
	// remainingDeck.print()
	cards.shuffle()
	cards.print()
}
