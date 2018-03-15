package main

func main() {
	cards := newDeck()

	cards.saveToFile("file.txt")
}
