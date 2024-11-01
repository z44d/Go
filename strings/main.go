package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello"
	b := "World"
	c := "!!"

	fmt.Println(a + b + c)
	// HelloWorld!!

	fmt.Println(a, b, c)
	// Hello World !!

	fmt.Println(fmt.Sprintf("%v %v %v", a, b, c))
	// Hello World !!

	fmt.Println(strings.Join([]string{a, b, c}, " "))
	// Hello World !!
}
