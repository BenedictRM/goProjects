package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//create a custom type that implements Writer interface
type logWriter struct{}

//out of net package we get a network interface
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		//make a slice with 99999 array positions -- needed so Read function has space to fill
		// bs := make([]byte, 99999)
		//Because Body implements Reader interface, it has a read func
		// resp.Body.Read(bs)
		// fmt.Println(string(bs))

		//Alternate solution -- stdout implements writer interface, body implements the reader interface
		//io.Copy(Writer, Reader) -- copy whats read in to output with writer
		io.Copy(os.Stdout, resp.Body)

		//Custom implementation:
		lw := logWriter{}

		io.Copy(lw, resp.Body)
	}
}

//just by defining this we now are technically implementing the Writer interface
//Implicitly implemented -- can cause issues
func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Printf("Just wrote in %v bytes", len(bs))
	return len(bs), nil
}
