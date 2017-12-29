package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 11)
	for number := range numbers {
		printNumber(number)
	}

	fmt.Println()
	fmt.Println("OR:")
	fmt.Println()

	index := 0
	for index <= 10 { // 10 <= 10 gives true; 11 <= 10 gives false
		printNumber(index)
		index++
	}
}

func printNumber(number int) {
	desc := "odd"

	if number%2 == 0 {
		desc = "even"
	}

	fmt.Printf("%2.1d is %s \n", number, desc)
}
