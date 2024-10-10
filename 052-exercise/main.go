package main

import "fmt"

var x = 40

const y = 41

func main() {
	z := 42
	fmt.Printf("The value of z is %vd and the type of z is %Td", z, z)
	fmt.Printf("The value of z is %v and the type of z is %T", z, z)
	fmt.Printf("The value of z is %v and the type of z is %T", z, z)
}
