/* WAP to print Fibonacci series with recursion */

package main

import "fmt"

func fibonacciWithRecursion(n1 int, n2 int, count int) {

	if count > 0 {
		n3 := n1 + n2
		fmt.Print(" ", n3)
		n1 = n2
		n2 = n3
		fibonacciWithRecursion(n1, n2, count-1)
	}
}

func main() {
	n1 := 0
	n2 := 1
	fmt.Print(n1, " ", n2)
	fibonacciWithRecursion(n1, n2, 4)
}
