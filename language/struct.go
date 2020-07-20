package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag bool
	counter int16
	pi float32
}

func main() {
	// -----------------------
	// Declare and initialize
	// -----------------------

	// Declare a variable of type example set to its zero value.
	// How much memory do we allocate for example?
	// a bool is 1 byte, int16 is 2 bytes, float32 is 4 bytes
	// Putting together, we have 7 bytes. However, the actual answer is 8.
	// That leads us to a new concept of padding and alignment.
	// The padding byte is sitting between the bool and the int16. The reason is because of alignment.

	// The idea of alignment: It is more efficient for hardware to read memory on its alignment boundary
	// we will take care of the alignment boundary issues so the hardware don't

	// Rule 1:
	// Depending on the size a particular value, Go determines the alignment we need. Every 2 bytes
	// value must follow a 2 bytes boundary. Since the bool value is only 1 byte and start a address 0,
	// than the  next int16 must start on address 2. The byte at address that get skipped over becomes a 1 byte padding.
	// Similarly, if it is a 4 bytes value then we will have a 3 bytes padding value

	var e1 example
	fmt.Printf("%+v\n", e1)

	// Rule 2:
	// The largest field represents the padding for entire struct.
	// we need to minimize the amount of padding as possible. Always lay out the field
	// from highest to smallest. This will push any padding down to bottom.
	//	type example1 struct {
	//		counter int64
	//		pi float32
	//		flag bool
	//	}

	// Declare a variable of type example and init using a struct literal.
	// Every line must end with coma.
	e2 := example{
		pi: 3.1415926,
		counter: 10,
		flag: true,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	// Declare a variable of an anonymous type and init using a struct literal.
	e3 := struct {
		flag bool
		counter int16
		pi float32
	}{
		counter: 10,
		pi: 3.1415926,
		flag: true,
	}

	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)

	// ----------------------------
	// Name type vs anonymous type
	// ----------------------------

	// If we have two name type identical struct, we can't assign one to another
	// For example, example1 and example2 are identical struct, var ex1 example1, var ex2 example2.
	// ex1 = ex2 is not allowed. we have to explicitly say that ex1 = example(ex2) by performing a conversion.
	// However, if ex is a value of identical anonymous struct type (like e3 above), then it is possible to
	// asign ex1 = ex
	var ex4 example
	ex4 = e3
	fmt.Printf("%+v\n", ex4)

	// The order of the fields in a struct should be consistent, otherwise it is not the same struct
	s1 := struct {
		num int64
		flag bool
	}{
		num: 1,
		flag: true,
	}

	s2 := struct {
		flag bool
		num int64
	}{
		num: 2,
		flag: false,
	}

	// cannot use s2 (type struct { flag bool; num int64 }) as type struct { num int64; flag bool } in assignment
	//s1 = s2
	fmt.Printf("%+v", s1)
	fmt.Printf("%+v", s2)

}