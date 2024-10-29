package main

import "fmt"

func main() {
	var myVar = "This is Var1"
	fmt.Println(myVar)

	var myVar2 string
	myVar2 = "This is var2"
	fmt.Println(myVar2)

	myVar3 := myVar2 + " This is var3"
	fmt.Println(myVar3)

	foo, bar := "Zaid", "Yazan"
	fmt.Println(foo, bar)

	var x, y = "Mahdi", "Sadiq"
	fmt.Println(x, y)

	var a, b = 6, 9
	fmt.Println(a, b)

	c := 2000
	fmt.Println(c)

	var myFloat float32 = 1.1
	fmt.Println(myFloat)

	myBool := true
	fmt.Println(myBool)

	mySlice := []string{"A", "B"}
	fmt.Println((mySlice))
}
