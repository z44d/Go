package main

import "fmt"

func main() {
	// Boolean / bool
	// true - false
	// 1 - 0
	myBool := true
	// fmt.Println(myBool)

	fmt.Printf("%v %T\n\n", myBool, myBool)

	// Integer / int
	// 1 2 3 4 5 6 7 8 9 0
	myInt := 1 + 3
	fmt.Printf("%v %T\n\n", myInt, myInt)

	// String / str
	// "a" "b" \ "a" + "b"

	myStr := "my" + " str"
	fmt.Printf("%v %T\n\n", myStr, myStr)

	// Floating / float32 / float64
	// 6.9 1.2 4.5
	myFloat := 6.9
	fmt.Printf("%v %T\n\n", myFloat, myFloat)

	// []string = Array, Slice
	var myString = []string{"a", "b"}
	fmt.Printf("%v %T\n\n", myString, myString)

}
