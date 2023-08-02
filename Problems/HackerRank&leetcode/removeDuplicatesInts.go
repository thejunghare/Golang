package main

import "fmt"

func removeDuplicates(nums []int) int {
/* 
	// Decalrig a empty map -> used to store the unqiue values from the slices
	bucket := make(map[int]bool)
	// This will store the dedeuplicate elements after removing duplicates from slice
	var result []int

	for _, val := range nums {
		if _, ok := bucket[val]; !ok { // add the element in map if not present already
			bucket[val] = true
			result = append(result, val)
		}
	}

	if len(nums) == 0 {
		return 0
	}

	return len(result) */

	i := 0
	   for j := 1; j < len(nums); j++ {
	       if nums[i] != nums[j] {
	           i++
	           nums[i] = nums[j]
	       }
	   }
	   return i + 1
}

func main() {
	// array := []int{1, 1, 2}
	array := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(array))
}
