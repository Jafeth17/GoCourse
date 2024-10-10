package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "Yona",
		last:    "Fonseca",
		flavors: []string{"Pistacho", "chocolate"},
	}

	p2 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"Vanilla", "Banana"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.flavors)
	fmt.Println(p2.flavors)

	for _, v := range p1.flavors {
		fmt.Println(p1.first, "favorite is", v)
	}
	for _, v := range p2.flavors {
		fmt.Println(p2.first, "favorite is", v)
	}
}
