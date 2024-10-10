package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}

	age := m["James"]
	fmt.Println("The age of James", age)

	age2 := m["Q"]
	fmt.Println("The age of Q ", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, abd here is the zero value of an int", v)
	}
}
