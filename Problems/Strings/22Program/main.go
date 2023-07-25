/* ▶ Golang code to find the index of last occurrence of specified substring in specified string. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World World"
	subString := "World"
	fmt.Println(findLastIndexOf(str, subString))
}

func findLastIndexOf(str string, subString string) int {
	return strings.LastIndex(str, subString)
}
