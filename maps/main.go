package main

import "fmt"

func main() {
	//declarations
	// var colors map[string]string

	// colors2 := make(map[string]string)
	// colors2["white"] = "#ffffff"

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#32CD32",
		"white": "ffffff",
	}
	printMap(colors3)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex: %s Color: %s\n", hex, color)
	}

}
