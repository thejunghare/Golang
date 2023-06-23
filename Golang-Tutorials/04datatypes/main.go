package main

import "fmt"

func main() {
	/*
		There are 3 basic data types:
			1. bool -> boolean value (true, false)
			2. Numeric -> integer value (1, 2.0)
			3. string -> string value ("Golang")
	*/

	// 1. boolean value (true, false)

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         //  untyped declaration with initial value

	fmt.Println(b1) // true
	fmt.Println(b2) // true
	fmt.Println(b3) // false
	fmt.Println(b4) // true

	// 2. Numeric -> integer value (1)
	var x int = 500  // Signed Integers -> both positive and negative values
	var y uint = 500 // Unsigned Integers -> only non-negative values

	fmt.Println(x) // 500
	fmt.Println(y) // 500

	// 2. Numeric -> floating value (2.0)
	var f1 float32 = 3.14
	var f2 float64 = 1.7e+308

	fmt.Println(f1) // 500
	fmt.Println(f2) // 500

	// 2. String -> String value ("Text")
	str := "Hello"
	fmt.Println(str) // Hello
}
