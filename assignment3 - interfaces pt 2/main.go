package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	validateArgs(os.Args)
	//Returns a File pointer and an err -- File implements the Reader interface
	file, _ := os.Open(os.Args[1])
	io.Copy(os.Stdout, file)
	//close file
	err := file.Close()
	if err != nil {
		fmt.Println("Got Error: ", err)
	}
}

func validateArgs(args []string) {
	if len(args) < 2 {
		fmt.Println("Missing command line argment: filename")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("Too many command line arguments.")
		os.Exit(1)
	}
}
