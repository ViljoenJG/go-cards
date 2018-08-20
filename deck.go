package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
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

func deckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bytes), ","))
}
