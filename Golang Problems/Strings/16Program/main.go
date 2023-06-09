/* ▶ Trimming space from both ends of string using TrimSpace() function in Golang ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "!!Hello World%%"
	fmt.Println("Trim space from both:",trimSpace(str))
	fmt.Println("Trim space from left:",trimSpaceLeft(str))
	fmt.Print("Trim space from right:",trimSpaceRight(str))
}

func trimSpace(str string) string {
	return strings.Trim(str, "!%")
}

func trimSpaceLeft(str string) string {
	return strings.Trim(str, "!")
}

func trimSpaceRight(str string) string {
	return strings.Trim(str, "%")
}
