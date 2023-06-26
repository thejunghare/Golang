package main

import "fmt"

func main() {
	/*
		1. with var keyword
			var array_name = [length]datatype{values} // here length is defined
			var array_name = [...]datatype{values} // here length is inferred

		2. with warlus operator
			array_name := [length]datatype{values}
			array_name := [...]datatype{values}
	*/

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
