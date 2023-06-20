/* WAP to print second largest number in array */

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{12, 35, 1, 10, 34, 1, 35}
	arr2 := []int{64, 34, 77, 43, 21}


	fmt.Println(secondLargestUsingSort(arr))       // with sort function
	fmt.Print(secondLargestWithoutUsingSort(arr2)) // without sort function
}

func secondLargestUsingSort(arr []int) int {
	sort.Ints(arr)
	return arr[(len(arr) - 2)]
}

func secondLargestWithoutUsingSort(arr []int) int {
	large_one := 0
	large_two := 0

	large_one = arr[0]

	for index := 1; index < len(arr); index++ {
		if large_one < arr[index] {
			large_two = large_one
			large_one = arr[index]
		} else if large_two < arr[index] {
			large_two = arr[index]
		}
	}

	return large_two
}
