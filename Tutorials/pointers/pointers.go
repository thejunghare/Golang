package main

import "fmt"

func main() {
	// declaring and initalizing an variable for storing hexadecimal value
	num := 0xFF

	// declaring and initalizing an pointer variable
	var num_address *int = &num
	fmt.Println("Address of variable num:", num_address)

	// uninitalized pointers
	fmt.Println("<-- uninitalized pointers -->")
	var uninitalized_pointer *int
	fmt.Println("Adderess of uninitalized pointer:", uninitalized_pointer)

	// using := syntax
	fmt.Println("<-- using := syntax -->")
	str := "hello"
	pointer := &str
	fmt.Println("Accessing memory address of str variable",pointer)
	fmt.Println("Accessing value of str variable with dereferencing operator",*pointer)

	// using var keyword with type interence
	fmt.Println("<-- using var keyword with type interence -->")
	var greet = "hello"
	var greet_pointer = &greet
	fmt.Println("Accessing memory address of greet_pointer variable",greet_pointer)
	fmt.Print("Accessing value of greet_pointer variable with dereferencing operator",*greet_pointer)
}