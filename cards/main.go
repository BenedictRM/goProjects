package main

func main() {
	cards := deck{newCard(), "Ace"}
	cards = append(cards, "New CARRRD")
	cards.print()
}

func newCard() string {
	return "5 of Diamonds"
}
