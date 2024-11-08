package main

import (
	"fmt"
	"strconv"
)

// Int

// int8 (1 Byte):

//     Range: -128 to 127
//     Can store small whole numbers, both positive and negative, in a single byte.

// int16 (2 Bytes):

//     Range: -32,768 to 32,767
//     Uses two bytes, allowing it to hold a larger range of positive and negative values.

// int32 (4 Bytes):

//     Range: -2,147,483,648 to 2,147,483,647
//     Uses four bytes, providing a very large range of integer values, common for larger calculations.

// int64 (8 Bytes):

//     Range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
//     With eight bytes, this data type can represent extremely large positive and negative numbers,
// 	   suitable for high-precision calculations or extensive data processing.

// Uint

// uint8 (1 Byte):

//     Range: 0 to 255
//     As it’s unsigned, it only stores non-negative values and can represent slightly larger positive values than int8 within the same storage size.

// uint16 (2 Bytes):

//     Range: 0 to 65,535
//     Holds larger positive values than int16 within two bytes, suitable for cases where only non-negative numbers are needed.

// uint32 (4 Bytes):

//     Range: 0 to 4,294,967,295
//     Provides a vast range of positive values using four bytes,
//	   often used in applications needing large positive numbers without any negatives.

// uint64 (8 Bytes):

//     Range: 0 to 18,446,744,073,709,551,615
//     Allows for extremely large positive values with eight bytes,
//	   ideal for high-precision calculations that don’t involve negative numbers.

// Float

// 1. float32

//     Size: 32 bits
//     Precision: 6-7 decimal digits
//     Range: ±1.18×10⁻³⁸ to ±3.4×10³⁸
//     Usage:
//	   Useful when you need to save memory or don't require high precision (like in some graphics applications or less precise scientific calculations).

// 2. float64

//     Size: 64 bits
//     Precision: 15-16 decimal digits
//     Range: ±2.23×10⁻³⁰⁸ to ±1.8×10³⁰⁸
//     Usage: The default floating-point type in Go, suitable for most applications requiring high precision,
//		such as scientific computations and financial calculations.

func main() {
	// int int8 int16 int32 int64

	var (
		a int8
		b int8
		c int16
		d int16
		e int32
		f int32
		g int64
		h int64
	)

	a = -128
	b = 127
	c = -32768
	d = 32767
	e = -2147483648
	f = 2147483647
	g = -9223372036854775808
	h = 9223372036854775807

	fmt.Println(
		fmt.Sprintf(
			"int8 : %v - %v\n"+"int16: %v - %v\n"+"int32: %v - %v\n"+"int64: %v - %v",
			a, b, c, d, e, f, g, h))

	// converting
	fmt.Println(int8(h), int16(e), int32(d), int64(a))

	// or just use int
	var x int = 1000000000000000000
	fmt.Println(x)

	// float32 float64

	// int => float32
	n := 5
	fmt.Println(fmt.Sprintf("%v %T", float32(n), float32(n)))

	// float64 => int
	z := 5.5
	fmt.Println(fmt.Sprintf("%v %T", int8(z), int8(z)))

	// int => str
	i := 15
	ii := strconv.Itoa(i)
	fmt.Println(fmt.Sprintf("%v %T", ii, ii)) // Outputs "15 string"

	// str => int
	s := "15"                               // Ensure `s` is purely numeric
	fmt.Println(fmt.Sprintf("%v %T", s, s)) // Outputs "15 string"

	if ss, err := strconv.Atoi(s); err == nil {
		fmt.Println(fmt.Sprintf("%v %T", ss, ss))
	}
}
