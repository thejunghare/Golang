/*
	Given a number N, print out all the numbers from 1 to N but when a number is divisible by 3 print out "Fizz" instead of the number, and when a number is divisible by 5 print out "Buzz" instead of the number. For numbers divisible by both 3 and 5 print "Fizz Buzz".
*/

package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}

func FizzBuzzUsingSwitch(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	n := 100
	FizzBuzz(n)
	FizzBuzzUsingSwitch(n)
}
