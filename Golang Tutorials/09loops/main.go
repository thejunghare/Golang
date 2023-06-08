package main

import "fmt"

func main() {
	/* loops */
	fmt.Println("Loops")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/* The continue statement */
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // This will skip 3
		}
		fmt.Println(i)
	}

	/* The break statement */
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // This will end the loop if i == 3
		}
		fmt.Println(i)
	}

	/* The Range Keyword */
	fmt.Println("Printing array elements using range keyword")
	arr := [...]string{"Google", "Apple", "Samsung"}
	for _, value := range arr { // this will print values
		fmt.Println(value)
	}

	fmt.Println("Printing array elements index using range keyword")
	for index, _ := range arr {
		fmt.Println(index)
	}

	/* Looping through array using for loop */
	fmt.Println("Looping through array using for loop")
	for index := 0; index < len(arr); index++ {
		fmt.Println(index)
	}
}
