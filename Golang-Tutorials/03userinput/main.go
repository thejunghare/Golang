package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	input, _ := reader.ReadString('\n')
	fmt.Print("You entered ", input)
	fmt.Printf("input is type of %T", input)
}
