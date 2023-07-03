package main

import (
	"fmt"
)

func lonelyInteger(a []int32) int32 {
	counts := make(map[int32]int)

	for _, num := range a {
		counts[num]++
	}

	for num, count := range counts {
		if count == 1 {
			return num
		}
	}

	return 0

	/*
		lonelyinteger := 0

		for _, num := range a {
			lonelyinteger ^= num
		}

		return lonelyinteger
	*/

	/*
		sort.Ints(a)
		for i := 0; i < len(a); i++ {
			for _, val := range a {
				if val != a[i] {
					lonelyinteger = a[i]
				}
			}
		}
	*/
}

func main() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Print(lonelyInteger(a))
}
