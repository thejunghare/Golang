/* 
	Given a list of numbers, determine if a specific number is in that list.
*/

package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	num := 5
	fmt.Print(isPresent(list, num))
}

func isPresent(list []int, num int) bool {
	for _, val := range list {
		if val == num {
			return true
		}
	}

	return false
}
