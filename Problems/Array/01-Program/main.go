/* WAP to sort array of integers */
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3}
	sort.Ints(arr)
	fmt.Print(arr)
}
