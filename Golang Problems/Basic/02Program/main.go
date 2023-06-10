/* GO Program to Check Armstrong Number */

package main

import "fmt"

func main() {

	var temp, remainder, result int
	number := 153
	temp = number

	for {
		remainder = temp % 10
		result += remainder * remainder * remainder
		temp /= 10

		if temp == 0 {
			break
		}
	}

	if result == number {
		fmt.Println("Armstrong Number")
	} else {
		fmt.Println("Not an Armstrong number")
	}
}
