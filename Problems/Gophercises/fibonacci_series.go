package main

import "fmt"

func FibonacciSeries(n int) {
	num1, num2, num3 := 0, 1, 0
	fmt.Print(num1, " ", num2)
	for i := 2; i < n; i++ {
		num3 = num1 + num2
		fmt.Print(" ", num3)
		num1 = num2
		num2 = num3
	}
}

func main() {
	FibonacciSeries(10)
}
