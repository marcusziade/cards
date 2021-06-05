package main

func main() {
	// cards := newDeckFromFile("my_cards")
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
