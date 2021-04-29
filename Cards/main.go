package main

func main() {
	// var card string = "Ace of Spades"
	//card := newCard() // Initialization
	//cards := deck{"Ace of Diamods", newCard()}
	//cards = append(cards, "Six of Spades")

	//cards.print()
	//printState()

	//cards := newDeck()

	//hand, remaining := deal(cards, 5)
	//cards[0] = "Hey"
	// cards.print()
	// hand.print()
	// remaining.print()

	//greeting := "Hi There"
	//fmt.Println([]byte(greeting))
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
