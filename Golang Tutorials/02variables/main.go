package main

import "fmt"

var variableOutSideTheMainFunction bool = true

// isTrue := true --> This wont work

const CONSTVARIABLE bool = true // public, Typed Const

const PI = 3.14 // Untyped Const

func main() {
	fmt.Println("Variables")

	/*  Multi-word variable names
	1. Camel Case
		-> variableName := value
	2. Pascal Case
		-> VariableName := value
	3. Snake Case
		-> variable_name := value
	*/

	var username string = "nimap"
	fmt.Println(username) // Print variable name

	fmt.Printf("username is of type: %T \n", username) // Print type of variable

	var anotherVariable = "loremcodes"
	fmt.Printf("anotherVariable is of type: %T \n", anotherVariable)

	isLoggedIn := true
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	fmt.Printf("variableOutSideTheMainFunction is of type: %T \n", variableOutSideTheMainFunction)

	// Multiple variables
	var a, b = 6, "Hello"
	// OR
	c, d := 7, "World"
	fmt.Println(a, b, c, d)
}
