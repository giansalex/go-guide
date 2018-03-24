package main

import "fmt"

type person struct {
	firstName string
	lastnNme  string
}

func main() {
	gian := person{firstName: "Giancarlos", lastnNme: "Salas"}

	fmt.Println(gian)
}
