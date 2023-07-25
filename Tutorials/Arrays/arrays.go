package main

import "fmt"

func main() {

	my_arr := [...]int{1, 2, 3, 4, 5}

	// Print array
	fmt.Println(my_arr)

	// Access array element
	fmt.Println(my_arr[4])

	// change values of element
	my_arr[4] = 2
	fmt.Println(my_arr[4])

	// Length of an array
	fmt.Println("length of my_array:", (len(my_arr)))

}
