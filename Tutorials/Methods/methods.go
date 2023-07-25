package main

import "fmt"

// define type Author
type Author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Method Receiver of author type
func (a Author) show() {
	fmt.Println("Result of struct type:", a.name)
}

// define non-struct type
type data int

// Defining a method with non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

// Method with pointer receiver
func (a *Author) showinfo(abranch string) {
	(*a).branch = abranch
}

// Main function
func main() {
	// Declaring and initializing value of Author type
	result := Author{
		name: "John",
		branch: "CS",
	}

	// calling the method
	result.show()

	// Calling the non-struct method
	val1 := data(5)
	val2 := data(2)
	res := val1.multiply(val2)
	fmt.Println("Result of non-struct type:", res)

	// Creating a pointer
	pointer := &result
	pointer.showinfo("CSE")

	fmt.Print(result.branch)
}
