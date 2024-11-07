package main

import (
	"fmt"
	"slices"
)

func main() {
	mySlice := []string{"Zaid", "Samer"}
	fmt.Println(mySlice)

	// append
	mySlice = append(mySlice, "Ballor")
	fmt.Println(mySlice)

	// delete
	mySlice = slices.Delete(mySlice, 0, 1) // delete first element
	fmt.Println(mySlice)

	// set last item
	mySlice[len(mySlice)-1] = "unknow"
	fmt.Println(mySlice)

	// for loop
	for x, y := range mySlice {
		// x is the index, y, is the value
		fmt.Println(fmt.Sprintf("%v - %v", x+1, y))
	}
}
