/* GO Program to take user input and addition of two strings */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print(takeUserInput())
}

func takeUserInput() string {
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter frist string")
	input1, _ := reader1.ReadString('\n')
	// fmt.Println(input1)

	fmt.Println("Enter second string")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	// fmt.Print(input2)

	return input1 + input2
}
