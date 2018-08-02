package main

import (
	"fmt"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
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
