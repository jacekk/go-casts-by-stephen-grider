package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 11)

	for number := range numbers {
		desc := "odd"
		if number%2 == 0 {
			desc = "even"
		}
		fmt.Printf("%2.1d is %s \n", number, desc)
	}
}
