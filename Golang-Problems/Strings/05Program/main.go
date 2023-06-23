/* ▶ Golang program to get the index of a specified character in the string  ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	subString := "W"

	fmt.Print(getIndexOfCharacter(str, subString))
}

func getIndexOfCharacter(str string, subString string) int {
	return strings.Index(str, subString)
}
