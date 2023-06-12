/* WAP to print Fibonacci series without recursion. */
package main

import "fmt"

func fibonacciWithOutRecursion(count int) {
	number1 := 0
	number2 := 1

	fmt.Print(number1, " ", number2)

	for index := 2; index <= count; index++ {
		number3 := number1 + number2
		fmt.Print(" ", number3)
		number1 = number2
		number2 = number3
	}
}

func main() {
	fibonacciWithOutRecursion(4)
}
