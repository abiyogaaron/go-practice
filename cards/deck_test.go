package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length is 52 cards, but deck length is %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Club" {
		t.Errorf("Expected last card is King of Club, but got %v", d[len(d)-1])
	}
}

func TestSetAndGetDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in the deck but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
