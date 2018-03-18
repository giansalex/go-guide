package main

func main() {
	cards := newDeck()

	hand, remaining := deal(cards, 2)

	hand.print()
	remaining.print()
}
