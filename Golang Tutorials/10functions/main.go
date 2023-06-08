package main

import "fmt"

/* Syntax for function without parameters
func functionName (){
	// code
}
*/

func functionWithoutParameters() {
	fmt.Println("functionWithoutParameters was called")
}

/* Syntax for function with parameters
func functionName (param1 type, param2 type){
	// code
}
*/

/* Function with single parameters */
func functionWithParameters(fname string) {
	fmt.Println("Hello", fname)
}

/* Function with multiple parameters */
func functionWithMultipleParameters(fname string, age int) {
	fmt.Println("My name is", fname, "and my age is", age)
}

/* Function Return */
func returnFunction(x int, y int) int {
	return x + y
}

/* Multiple return function */
func multipleReturnFunction(x int, y string)(result int, txt string){
	result = x + x
	txt = y + "world"
	return
}

/* Main function */

func main() {
	functionWithoutParameters() /* Calling function without parameters */

	functionWithParameters("Golang") /* Calling function with parameters */

	functionWithMultipleParameters("Paddy", 22) /* Calling function with parameters */

	fmt.Println(returnFunction(2, 3)) /* Calling return function with parameters */

	_,b := multipleReturnFunction(5, "hello ")  /* Calling return function with multiple return function */

	fmt.Println(b)
}
