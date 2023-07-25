/*
Find two numbers in a list that sum to a given amount.

FindTwoThatSum will look for two numbers in the provided numbers list that sum up to the sum argument. It will then return the indices of each of these numbers.

If there are multiple correct answers, any of them may be used. If there are no correct answers, (-1, -1) should be returned.
*/

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
	for i, val1 := range list {
		for j, val2 := range list {
			if i == j {
				continue
			}
			if val1+val2 == sum {
				return i, j
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
