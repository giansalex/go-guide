package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}
