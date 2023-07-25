package main

import (
	"fmt"
)

func findMedian(arr []int32) int32 {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	length := int32(len(arr))

	median := int32(0)

	for _, val := range arr {
		median += val
	}

	return median / length

}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int32, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	fmt.Print(findMedian(arr))
}
