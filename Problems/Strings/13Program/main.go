/* ▶ Golang code to check the specified substring is the suffix of the given string. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(totalOccurence("hello"))
}

func totalOccurence(str string) int {
	return strings.Count(str, "l")
}
