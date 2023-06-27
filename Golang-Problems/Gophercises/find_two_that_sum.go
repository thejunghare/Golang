package main

import "fmt"

func FindTwoThatSum(list []int, sum int) []int {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == sum {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func FindTwoThatSumUsingRangeKey(list []int, sum int) (int, int) {
	for i,val1 := range list {
		for j, val2 := range list {
			if i == j {
				continue
			}
			if val1 + val2 == sum {
				return i , j
			}
		}
	}
	
	return -1, -1
}

func main() {
	list := []int{2, 3, 5, 7}
	sum := 9

	fmt.Print(FindTwoThatSumUsingRangeKey(list, sum))
}
