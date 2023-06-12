package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{12, 35, 1, 10, 34, 1, 35}
	sort.Ints(arr)

	fmt.Println(arr[len(arr)-3])

	fmt.Println(largestElement(12, 35, 1, 10, 34, 1, 35))
}

func largestElement(arr ...int) int {
	largest := arr[0]

	for index := 0; index < len(arr); index++ {
		if arr[index] > largest {
			largest  = arr[index]
		}
	}

	return largest
}
