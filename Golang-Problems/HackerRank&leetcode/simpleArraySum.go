/* Given an array of integers, find the sum of its elements. */

package main

import "fmt"

func main() {
	ar := []int32{1, 2, 3, 4}
	fmt.Print(simpleArraySum(ar))
}

func simpleArraySum(ar []int32) int32 {
	sum := 0
	for _, val := range ar {
		sum += int(val)
	}
	return int32(sum)
}
