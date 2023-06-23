/* ▶Golang program to check a specified string contains a Unicode code point. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "A"
	unicode := 64.6
	fmt.Println(checkIfContain(str, unicode))
}

func checkIfContain(str string, unicode float64) bool {
	return strings.ContainsRune(str, rune(unicode))
}
	