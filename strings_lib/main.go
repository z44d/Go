package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b := "string1", "string2"
	// The result is -1 because "string1" is lexicographically less than "string2".
	// Lexicographical comparison is similar to alphabetical order.
	// Since "string1" comes before "string2" in this order, the result is -1.
	fmt.Println(strings.Compare(a, b))

	// The result is 'true' because the variable a contains "1"
	fmt.Println(strings.Contains(a, "1")) // true

	// The result is 'true' bacause the variable b contains 'g'
	fmt.Println(strings.ContainsAny(b, "abcdefg")) //true
}
