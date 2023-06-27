/* WAP to reverse an integer without converting it to a string, without using any builtin methods */

package main

import "fmt"

func main() {
	number := 23
	reverse := 0

	for number > 0 {
		remainder := number % 10
		reverse = (reverse * 10) + remainder
		number /= 10
	}

	fmt.Println("Reversed number is:", reverse)
}
