package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	fmt.Print(cards.toString())
}
