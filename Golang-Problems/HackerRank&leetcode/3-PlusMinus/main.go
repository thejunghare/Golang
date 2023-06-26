package main

import "fmt"

func main() {
	negative_count := 0
	positive_count := 0
	zero_count := 0

	for _, val := range arr {
		if val < 0 {
			negative_count++
		} else if val > 0 {
			positive_count++
		} else {
			zero_count++

		}
	}

	// Length of array
	length_array := len(arr)

	// Ratio of Positive
	fmt.Println(float64(positive_count) / float64(length_array))

	// Ratio of Negative
	fmt.Println(float64(negative_count) / float64(length_array))

	// Ratio of zero
	fmt.Println(float64(zero_count) / float64(length_array))
}
