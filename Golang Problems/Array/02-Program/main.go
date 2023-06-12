/* WAP to sort array of strings */
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"q", "w", "e", "r", "t", "y"}
	sort.Strings(arr)
	fmt.Print(arr)
}
