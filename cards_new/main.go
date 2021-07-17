package main

func main() {
	cards := newDeckFromFile("my_cards")
	// cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
	// fmt.Print([]byte(cards.toString()))
	// hand, remainingCards := deal(cards, 5)
	// cards.print()
	// hand.print()
	// remainingCards.print()
}

func newCard() string {
	return "Ace of Spades"
}
