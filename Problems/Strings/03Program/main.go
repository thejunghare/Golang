/* Golang code to convert specified string in uppercase */

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello"
	fmt.Print(toUpperCase(str))
}

func toUpperCase(str string) string {
	return strings.ToUpper(str)
}
