package main

import "fmt"

func variadicFunction(a int, b ...int) {
	fmt.Println(a)

	for _, index := range b {
		fmt.Println(index)
	}
}

func main() {
	fmt.Println("<- passing one argument to b ->")
	variadicFunction(1, 2) // passing one argument to b

	fmt.Println("<- passing multiple argument to b ->")
	variadicFunction(1, 2, 3, 4, 5) // passing multiple argument to b

	fmt.Println("<- passing zero argument to b ->")
	variadicFunction(1) // passing zero arguments to b
}
