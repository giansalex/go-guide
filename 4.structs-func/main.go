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

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
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

	gianPointer := &gian
	gianPointer.updateName("Alexander")
	gian.print()
}
