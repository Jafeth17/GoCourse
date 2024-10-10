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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println("--------MAP SECTION---------")
	for _, v := range m {
		//fmt.Println(v)
		for _, v2 := range v.flavors {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
