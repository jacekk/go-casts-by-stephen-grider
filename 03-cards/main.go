package main

import (
	"fmt"
)

func main() {
	// var card string = "Uno"
	// card := "Uno" // equivalent with upper line
	// card = "Dos" // changing already declared
	// card = newCard()

	filename := "my-cards.txt"

	cards := newDeck()
	subset := cards[2:5] // 3 elements
	leftDeck, rightDeck := deal(cards, 5)
	subset.saveToFile(filename)
	fromFile := newDeckFromFile(filename)

	fmt.Println("\nAll:")
	cards.print()
	fmt.Println("\nSubset:")
	subset.print()
	fmt.Println("\nRight:")
	leftDeck.print()
	fmt.Println("\nLeft:")
	rightDeck.print()
	fmt.Println("\nSubset from file:")
	fromFile.print()
	fmt.Println("\nRandomized subset:")
	subset.shuffle()
	subset.print()
}

func newCard() string { // can be defined after usage
	return "Tres"
}
