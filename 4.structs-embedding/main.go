package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastnNme  string
	contact   contactInfo
}

func main() {
	gian := person{
		firstName: "Giancarlos",
		lastnNme:  "Salas",
		contact: contactInfo{
			email:   "giansalex@gmail.com",
			zipCode: 9400,
		},
	}

	fmt.Println(gian)
	fmt.Printf("%+v", gian)
}
