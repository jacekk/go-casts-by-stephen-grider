package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	lastIndex := len(d) - 1

	expectedLength := 16
	firstExpected := "Ace of Spades"
	lastExpected := "Four of Clubs"

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of '%d', but got '%v'.", expectedLength, len(d))
	}

	if d[0] != firstExpected {
		t.Errorf("Expected first card of '%s', but got '%v'.", firstExpected, d[0])
	}

	if d[lastIndex] != lastExpected {
		t.Errorf("Expected last card of '%s', but got '%v'.", lastExpected, d[lastIndex])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	path := "_decktesting"

	os.Remove(path)
	deck := newDeck()
	deck.saveToFile(path)
	loadedDeck := newDeckFromFile(path)
	os.Remove(path)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got '%v'.", len(loadedDeck))
	}
}
