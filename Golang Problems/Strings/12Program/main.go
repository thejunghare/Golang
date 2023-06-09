/* ▶ Golang code to check the specified substring is the suffix of the given string. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	subString := "lo"
	fmt.Println(hasSuffix(str, subString))
}

func hasSuffix(str string, subString string) bool {
	return strings.HasSuffix(str, subString)
}