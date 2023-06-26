package main

import "fmt"

func main() {
	for index := 1; index <= 100; index++ {
		if index%3 == 0 && index%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if index%3 == 0 {
			fmt.Println("Fizz")
		} else if index%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(index)
		}
	}
}
