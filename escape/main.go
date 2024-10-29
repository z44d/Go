package main

import "fmt"

func main() {
	// \b - Back Space
	fmt.Println("Hello \bWorld")
	// Output: "HelloWorld"

	// \n - New Line
	fmt.Println("Hello \nWorld")
	// Output:
	// Hello
	// World
	// Or
	fmt.Println(`Hello
World
	`)

	// \r - return
	fmt.Println("AAAAA\rFuck")
	// Output: FuckA

	// \t - Tap
	// fmt.Println("	Hi!")
	fmt.Println("\tHi!")
	// Output: "	Hi!"

	// Egnoring back slash

	fmt.Println("Hello \\\\ World")
	// Output: Hello \\ World

	// Or
	fmt.Println(`Hello \\ World`)

	// Printing the Quotes
	fmt.Println("Hello \"World\"")
	fmt.Println(`Hello "World"`)
}
