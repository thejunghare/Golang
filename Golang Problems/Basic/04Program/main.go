/* GO Program to Find Factorial of a Number */

package main

import "fmt"

func main() {
	number := 3
	fmt.Print(calculateFactorail(number))
}

func calculateFactorail(number int) int {
	if number == 1 || number == 0 {
		return number
	}

	return number * calculateFactorail(number-1)
}
