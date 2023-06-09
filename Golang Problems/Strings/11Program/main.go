/* ▶ Golang code to check the specified substring is the prefix of the given string. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "india"
	subString := "in"
	fmt.Println(hasPrefix(str, subString))
}

func hasPrefix(str string, subString string) bool {
	return strings.HasPrefix(str, subString)
}
