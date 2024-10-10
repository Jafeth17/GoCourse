package main

import "fmt"

func main() {
	x := []int{42, 41, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Printf("%v \t %T \t %v\n", v, v, i)
	}

	fmt.Printf("%T\t %#v\n", x, x)
	fmt.Printf("%T\t %v\n", x, x)

	fmt.Println(x[:5])
	fmt.Println(x[6:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
