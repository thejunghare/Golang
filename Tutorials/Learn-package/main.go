package main

import (
	"fmt"
	"learnpackage/addTwoNumbers"
)

func main() {
	fmt.Println("Main function")

	fmt.Println("Addition of two numbers is ", addTwoNumbers.Calculate(1, 2))
}
