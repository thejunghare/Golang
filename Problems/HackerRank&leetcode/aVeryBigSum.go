/* In this challenge, you are required to calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large. */

package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	sum := int64(0)
	for _, val := range ar {
		sum += val
	}
	return sum
}

func main() {
	ar := []int64{1000000001,1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Print(aVeryBigSum(ar))
}
