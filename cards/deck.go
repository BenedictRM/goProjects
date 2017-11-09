package main

import( 
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

//Create a new type of Deck which is a slice of strings
//think of this as saying this new deck type extends the behaviour of a slice of strings -- but we don't use oop language in Go ;)
type deck []string

//generate a deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename)
	if(err != nil){
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//Type convert to a string
	return deck(strings.Split(string(bs), ","))
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
