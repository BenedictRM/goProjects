package main

func main() {
	cards := newDeckFromFile("MyDeckOfCards")
	// hand, remainingDeck := deal(cards, 5)
	cards.shuffle()
	cards.print()
}
