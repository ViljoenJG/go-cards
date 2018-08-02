package main

func main() {
	cards := newDeck()
	hand, cards := deal(cards, 5)
	cards.print()
	hand.print()

	hand.saveToFile("myHand")
}
