// This program demonstrates basic comparison operators in Go. It compares
// two integer values, A and B, using greater-than, less-than, and equality
// operators. The results of each comparison are printed to show whether
// the expressions evaluate as true or false.

package main

import "fmt"

func main() {
	A := 1
	B := 2
	fmt.Println(A > B) // true
	fmt.Println(B > A) // false
	fmt.Println(B < A) // true
	// u can use ()>\(<) + (=)
	// A >= B | B <= A

	fmt.Println(A == B)   // false
	fmt.Println(A == B-1) // true

}
