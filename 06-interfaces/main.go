package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello world!"
}

func (spanishBot) getGreeting() string {
	return "Holla!"
}

func printGreenting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreenting(eb)
	printGreenting(sb)
}
