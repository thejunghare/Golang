package main

import "fmt"

func main() {
	x, y, z := 21, 23, 24

	/* if statement */
	if x < y {
		fmt.Println("x is smaller")
	}

	/* if else statement */
	if x > y {
		fmt.Println("x is greater")
	} else {
		fmt.Println("x is smaller")
	}

	/* else if statement */
	if x < y {
		fmt.Println("y is greater")
	} else if x > y {
		fmt.Println("x is smaller")
	} else {
		fmt.Println("x is greater")
	}

	/* Nested if */
	if x < y {
		if x < z {
			fmt.Println("x is smallest")
		}
	}

}
