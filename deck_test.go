package main

import (
	"os"
	"testing"
)

func assertEqualInt(val int, expected int, t *testing.T) {
	if val != expected {
		t.Errorf("Expected value of %v, but got %v", expected, val)
	}
}

func assertEqualStr(val string, expected string, t *testing.T) {
	if val != expected {
		t.Errorf("Expected value of %v, but got %v", expected, val)
	}
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	assertEqualInt(len(d), 52, t)
	assertEqualStr(d[0], "Ace of Spades", t)
	assertEqualStr(d[len(d)-1], "King of Hearts", t)
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	filename := "_testdeck"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := deckFromFile(filename)

	assertEqualInt(len(loadedDeck), 52, t)
	assertEqualStr(loadedDeck[0], "Ace of Spades", t)
	assertEqualStr(loadedDeck[len(loadedDeck)-1], "King of Hearts", t)

	os.Remove(filename)
}
