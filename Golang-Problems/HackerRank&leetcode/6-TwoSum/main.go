package main

import "fmt"

func twosum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	target := 9
	nums := []int{15, 2, 11, 5, 7}
	fmt.Print(twosum(nums, target))
}
