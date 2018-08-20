package main

func main() {
	filename := "my_hand"
	cards := newDeck()
	cards.shuffle()
	hand, cards := deal(cards, 5)

	hand.saveToFile(filename)

	loaded := deckFromFile(filename)
	loaded.print()
}
