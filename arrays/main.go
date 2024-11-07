package main

import (
	"fmt"
)

func main() {
	// var arr [3]string = [3]string{"a", "b", "c"}
	var arr = [3]string{"a", "b", "c"}
	fmt.Println(arr)

	arr1 := [3]string{"a", "b", "c"}
	fmt.Println(arr1)
	new_arr := []string{}

	// del item
	for _, x := range arr1 {
		if x == "a" {
			continue
		} else {
			new_arr = append(new_arr, x)
		}
	}

	// or slices.Delete

	fmt.Println(new_arr)

	// edit value of item
	arr1[1] = "B"
	fmt.Println(arr1)
}
