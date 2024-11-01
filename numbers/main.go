package main

import "fmt"

func main() {
	// + | - | / | * | %

	// +
	A := 24
	A = A + 1
	fmt.Println(A) // 25

	// -
	B := 19
	B = B - 1
	fmt.Println(B) // 18

	// *
	C := 2
	C = C * 5
	fmt.Println(C) // 10

	// % Modulo
	// The modulo operator (%)
	// returns the remainder of integer division between two numbers.
	fmt.Println(8 % 2) // 0
	fmt.Println(9 % 2) // 1

}
