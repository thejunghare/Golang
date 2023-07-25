package main

import "fmt"

func myFunction() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	integer := myFunction()
	fmt.Println(integer())
	fmt.Println(integer())
	fmt.Println(integer())
}
