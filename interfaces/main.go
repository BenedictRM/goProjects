package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//interface implementation
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//NOTE: if not using receiver value, you can optionally omit it
func (englishBot) getGreeting() string {
	//assume custom logic for generating an eng greeting
	return "Hello"
}

func (spanishBot) getGreeting() string {
	//assume cust logic
	return "Hola"
}
