/* ▶ Golang program to demonstrate the strings.ContainsAny() function. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "india"
	subString := "ia"

	fmt.Println(doContain(str, subString))
}

func doContain(str string, subString string) bool {
	return strings.ContainsAny(str, subString)
}
