package main

import (
	"fmt"
	"net/http"
	"os"
)

//out of net package we get a network interface

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		fmt.Println(*resp)
	}
}
