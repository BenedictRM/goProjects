package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range numbers {
		fmt.Printf("%v is %s\n", v, evenOrOdd(v))
	}
}

func evenOrOdd(i int) string {
	answer := ""
	if i%2 == 0 {
		answer = "even"
	} else {
		answer = "odd"
	}
	return answer
}
