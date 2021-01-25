package main

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// // hand.print()
	// // remainingDeck.print()
	// greeting := "Hi there"
	cards.saveToFile("my_cards")
}

func newCard() string {
	return "Five of Diamonds"
}
