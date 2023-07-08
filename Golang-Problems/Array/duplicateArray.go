package main

import (
	"fmt"
)

func removeDuplicateElments(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, val := range nums {
		if !seen[val] {
			result = append(result, val)
			seen[val] = true
		}
	}

	return result
}

func main() {
	nums := []int{1, 4, 3, 5, 2, 1}

	fmt.Print(removeDuplicateElments(nums))
}
