/* 
	Given a list of numbers, sum them all up and return the sum.
*/

package main

import "fmt"

func sum(list []int) int {
	sum := 0

	for _, value := range list {
		sum += value
	}

	return sum
}

func sumUsingRecursion(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + sumUsingRecursion(list[1:])
}

func main() {
	list := []int{4, 2, 22, -10, 8}
	fmt.Print("Sum of list: ", sum(list))
}
