package main

import "fmt"

// Made by me
func input(a string) string {
	fmt.Printf(a)
	var value string
	fmt.Scanln(&value)
	return value
}

func main() {
	var name = input("Enter Your Name : ")
	fmt.Println()
	fmt.Printf("Your Name iss %v", name)
	fmt.Println()
}

// func main() {
// 	fmt.Printf("Enter Your Name : ")
// 	var name string
// 	fmt.Scanln(&name) // Add the value from user into "name" var.

// 	fmt.Println()

// 	fmt.Printf(
// 		"This is your name : %v",
// 		name)

// 	fmt.Println()
// }
