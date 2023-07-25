/* ▶  Golang program to demonstrate the strings.Compare() function. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	subString := "Hello"

	fmt.Println(compareString(str, subString))
}

func compareString(str string, subString string) int {
	return strings.Compare(str, subString)
}
