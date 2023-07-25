/* ▶ Golang code to change the specified string into title case. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "12 rules for life"
	fmt.Println(toTitleCase(str))
}

func toTitleCase(str string) string {
	return strings.Title(str)
}
