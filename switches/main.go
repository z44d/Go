package main

import (
	"flag"
	"fmt"
)

// flag.dataType
// flag.String
// flag.Int
// flag.Bool
// flag.Float

func main() {
	name := flag.String("name", "Zaid ballor", "Enter your full name")
	age := flag.Int("age", 17, "Enter your age")
	hasMarried := flag.Bool("married", false, "have you married?")
	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*hasMarried)
}
