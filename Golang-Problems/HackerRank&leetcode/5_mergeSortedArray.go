package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(sort.Ints(nums1[:n]), nums2...)
	sort.Ints(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := len(nums1)
	n := len(nums2)

	merge(nums1, nums2, m, n)
}
