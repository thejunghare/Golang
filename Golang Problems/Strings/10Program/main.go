/* ▶ Golang program to demonstrate the strings.Join() function. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"India", "is", "a", "great", "country"}

	fmt.Println(joinFunction(arr...))
}

func joinFunction(arr ...string) string {
	return strings.Join(arr, " ")
}
