package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename+".deck", []byte(d.toString()), 0666)
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, val := range values {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}
