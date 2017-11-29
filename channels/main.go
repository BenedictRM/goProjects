package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://amazon.com",
	}
	//Single Thread
	for _, l := range links {
		checkLink(l)
	}
	fmt.Println("*********************************")
	//create a channel for interprocess communication
	c := make(chan string)
	//Try running each request in it's own routine
	for _, l := range links {
		//continuously spawn new go routines
		go checkLinkGoRoutine(l, c)
	}
	//receive the channel updates and log it--this blocks main until a message is received
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(l string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println("ERROR: ", l, err)
	} else {
		fmt.Println(l, resp.Status)
	}
}

func checkLinkGoRoutine(l string, c chan string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println("ERROR: ", l, err)
		//pass update back to main
		c <- err.Error()
	} else {
		// fmt.Println(l, resp.Status)
		//pass update back to main
		c <- l + " " + resp.Status
	}
}
