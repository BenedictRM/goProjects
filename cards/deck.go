package main

import "fmt"

//Create a new type of Deck which is a slice of strings
//think of this as saying this new deck type extends the behaviour of a slice of strings -- but we don't use oop language in Go ;)
type deck []string

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}