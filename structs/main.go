package main

import "fmt"

type Id struct {
	name   string
	age    int
	skin   string
	height int
}

func main() {
	// person1 := Id{name: "Zaid", age: 17, skin: "nigga", height: 172}
	person1 := Id{
		"Zaid", 17, "nigga", 172}
	fmt.Println(person1)

	// edit element
	person1.age = 18

	// Get elements
	fmt.Println(
		fmt.Sprintf(
			"Name: %v\nAge: %v\nSkin: %v\nHeight: %v",
			person1.name,
			person1.age,
			person1.skin,
			person1.height))
}
