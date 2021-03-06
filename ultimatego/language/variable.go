package main

import "fmt"

func main() {
	// --------------
	// Built-in-types
	// --------------

	// Type provides integrity and readability
	// - What is the amount of memory that we allocate?
	// - what does that memory represent?

	// Type can be specific such as int32 or int64
	// For example,
	// - uint8 contains a base 10 number using one byte of memory. (1 byte = 32 bit)
	// - int32 contains a base 10 number using 4 byte of memory.

	// When we declare a type without being very specific, such as uint or int, it gets mapped
	// base on the architecture we are building the code against.
	// On a 64-bit architecture, the word size is 64 bit (8 bytes), address size is 64 bit
	// then our integer should be 64 bit.

	// -------------------
	// Zero value concept
	// -------------------

	// Every single value we create must be initialized. If we don't specify it, it will be set to the zero value
	// The entire allocation of memory, we reset that bit to 0.
	// - Boolean false
	// - Integer 0
	// - Floating point 0
	// - Complex 0i
	// - String "" (empty string)
	// - Pointer nil

	// - String are a series of uint8 types.
	// A string is two word data structure: first word represents a pointer to a backing array(https://stackoverflow.com/questions/31263591/what-is-array-backed-data-structure),
	// second word represents its length.
	// If it is a zero value them the first word is nil, the second word is 0.

	// -----------------------
	// Declare and initialize
	// -----------------------

	// var is the only guarantee to initialize a zero value for a type
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Using the short variable declaration operator, we can define and initialize at the same time.
	aa := 10
	bb := "hello"
	cc := 3.1415926
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.1415926 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// ----------------------
	// Conversion vs casting
	// ----------------------

	// Go doesn't have casting, but conversion.
	// Instead of telling a compiler to pretend to have some more bytes, we have to allocate more memory
	// (https://stackoverflow.com/questions/3166840/what-is-the-difference-between-casting-and-conversion)
	// Casting is a term describing syntax (hence the Syntactic meaning).
	// Conversion is a term describing what actions are actually taken behind the scenes (and thus the Semantic meaning).

	// Specify type and perform a conversion
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)

}
