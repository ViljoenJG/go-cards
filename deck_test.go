package main

import (
	"os"
	"testing"
)

type LocalT struct {
	t *testing.T
}

func (l LocalT) assertEqualInt(val int, expected int) {
	if val != expected {
		l.t.Errorf("Expected value of %v, but got %v", expected, val)
	}
}

func (l LocalT) assertEqualStr(val string, expected string) {
	if val != expected {
		l.t.Errorf("Expected value of %v, but got %v", expected, val)
	}
}

func TestNewDeck(t *testing.T) {
	lt := LocalT{t}
	d := newDeck()

	lt.assertEqualInt(len(d), 52)
	lt.assertEqualStr(d[0], "Ace of Spades")
	lt.assertEqualStr(d[len(d)-1], "King of Hearts")
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	lt := LocalT{t}
	filename := "_testdeck"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := deckFromFile(filename)

	lt.assertEqualInt(len(loadedDeck), 52)
	lt.assertEqualStr(loadedDeck[0], "Ace of Spades")
	lt.assertEqualStr(loadedDeck[len(loadedDeck)-1], "King of Hearts")

	os.Remove(filename)
}
