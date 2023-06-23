package main

import "fmt"

func main() {
	/*
		slice_name := []datatype{values}

		// with make() function
		slice_name := make([]type, length, Capacity)

	*/

	arr := [...]int{1, 2, 3, 4, 5} // This is an array
	slice := []int{1, 2}           // This is an slice
	fmt.Println(slice)

	// Create slice from an array
	new_slice := arr[2:(len(arr) - 1)]
	fmt.Println(new_slice)

	// Length of slice
	fmt.Println("Length of slice:", len(new_slice))

	// Capacity of slice
	fmt.Println("Capacity of slice:", cap(new_slice))

	second_slice := make([]int, 5, 10)
	fmt.Println("Elements of the second_slice:", second_slice)
	fmt.Println("Capacity of the second_slice:", cap(second_slice))
	fmt.Println("Length of the second_slice:", len(second_slice))

	// Access elements of slice
	fmt.Println("0 element of slice is:", second_slice[0])

	// Change elements of slice
	second_slice[1] = 2
	fmt.Println("1 element of slice is:", second_slice[1])

	// Append elements to slice
	second_slice = append(second_slice, 1, 2)
	fmt.Println(second_slice)

	// Append one slice to another
	third_slice := []int{3, 4}
	fourth_slice := append(second_slice, third_slice...)
	fmt.Println(fourth_slice)
}
