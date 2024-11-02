package main

import (
	"fmt"
)

func main() {
	name := "Zaid"
	age := 17

	// Printing with different methods
	fmt.Print("Name: ")
	fmt.Println(name)
	fmt.Printf("Age: %d\n", age)

	// Using Sprintf
	info := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(info)

	// Using Errorf to create a custom error
	err := fmt.Errorf("user %s is underage", name)
	fmt.Println(err)

	// Scanning input from user
	var inputName string
	var inputAge int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&inputName, &inputAge)
	fmt.Printf("Hello, %s. You are %d years old.\n", inputName, inputAge)
}
