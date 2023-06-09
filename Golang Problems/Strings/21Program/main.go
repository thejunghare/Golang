/* ▶ Golang program to get the index of a specified character in the specified string. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	byteValue := byte('r')

	fmt.Println(getIndexOfCharacter(str, byteValue))
}

func getIndexOfCharacter(str string, byteValue byte) int {
	return strings.IndexByte(str, byteValue)
}
