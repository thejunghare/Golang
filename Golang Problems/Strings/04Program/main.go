/* ▶ Golang program to convert specified string in lowercase ◀ */

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "HELLO"
	fmt.Print(toLowerCase(str))
}

func toLowerCase(str string) string {
	return strings.ToLower(str)
}
