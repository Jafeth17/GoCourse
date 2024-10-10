package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("x and y are %v and %v \t \n", x, y)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	// if x < 4 && y < 4 {
	// 	fmt.Println("both are less than four")
	// } else if x > 6 && y > 6 {
	// 	fmt.Println("Both are more than 6")
	// } else if x >= 4 && x <= 6 {
	// 	fmt.Println("X is from 4 to 6 inclusive on both numbers")
	// } else if y != 5 {
	// 	fmt.Println("y is not 5")
	// } else {
	// 	fmt.Println("None of the previous")
	// }

	// switch {
	// case x < 4 && y < 4:
	// 	fmt.Println("both are less than four")
	// case x > 6 && y > 6:
	// 	fmt.Println("Both are more than 6")
	// case x >= 4 && x <= 6:
	// 	fmt.Println("X is from 4 to 6 inclusive on both numbers")
	// case y != 5:
	// 	fmt.Println("y is not 5")
	// default:
	// 	fmt.Println("None of the previous were met")
	// }
}
