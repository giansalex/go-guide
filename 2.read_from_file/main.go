package main

func main() {
	cards := newDeckFromFile("file.txt")

	cards.print()
}
