package main

import "fmt"

func main() {
	foo()
	bar("Yona")
	s := aloha("Yona")
	fmt.Println(s)
	s1, n := dogYears("Yona", 23)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}
