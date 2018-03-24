package main

import "fmt"

type person struct {
	firstName string
	lastnNme  string
}

func main() {
	var gian person

	gian.firstName = "Giancarlos"
	gian.lastnNme = "Salas"

	fmt.Println(gian)
	fmt.Printf("%+v", gian)
}
