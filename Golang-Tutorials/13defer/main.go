/*
	defer -> defer means postpone

	1. defer is used to delay execution of a function until surrounding function completes.

	2. defer function are executed in last in first out order.
*/

package main

import "fmt"

func main() {
	defer myDeferFunction()

	defer fmt.Println("First")

	defer fmt.Println("Second")

	defer fmt.Println("Third")

	fmt.Println("Fourth")

	fmt.Println("Five")

	myFunction()

	/*
		Result would be displayed in following order
		Fourth
		Five
		function
		Third
		Second
		First
		defer function
	*/
}

func myFunction() {
	fmt.Println("function")
}

func myDeferFunction() {
	fmt.Println("defer function")
}
