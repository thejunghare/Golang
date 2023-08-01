package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(arr []int) []int {
	maps := make(map[int]bool)
	var newarr []int

	for _, val := range arr {
		if _, ok := maps[val]; !ok {
			maps[val] = true
			newarr = append(newarr, val)
		}
	}
	return newarr
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// remove the duplicates elements
	nums1 = removeDuplicates(nums1)
	nums2 = removeDuplicates(nums2)

	// Sort the arrays
	sort.Ints(nums1)
	sort.Ints(nums2)

	// Merge the arrays
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

func main() {
	arr := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr, 3, arr2, 3)
	fmt.Println( arr)
}
