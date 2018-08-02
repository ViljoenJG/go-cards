package main

func main() {
	filename := "my_hand"
	cards := newDeck()
	hand, cards := deal(cards, 5)

	hand.saveToFile(filename)

	loaded := deckFromFile(filename)
	loaded.print()
}
