package main

import "fmt"

func main() {
	x := "Hola"

	fmt.Println(x)
	x = "Adios!"
	// Changing the value
	fmt.Println(x)

	const y = "La La La"
	// y = "" // Unchangable.
	fmt.Println(y)
}
