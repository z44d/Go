package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"cat": 40,
		"dog": 1,
		"me":  1000 * 1000, // 1M
	}
	fmt.Println(myMap)

	// print an value of key
	fmt.Println(myMap["cat"])
	fmt.Println(myMap)

	// change value
	myMap["dog"] = 45

	// create empty map
	var mappy = make(map[string]any)
	mappy["cat"] = 5
	mappy["dog"] = 2.5
	mappy["me"] = 1000 * 1000
	fmt.Println(mappy)

	// check if element exist
	price, cat := myMap["cat"]
	fmt.Println(fmt.Sprintf("Cat is available? %v and the price is %v", cat, price))

	for element, value := range myMap {
		fmt.Println(fmt.Sprintf("Element %v, has value %v", element, value))
	}

	fmt.Println(len(myMap))
}
