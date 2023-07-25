/* ▶ Golang code to demonstrate the example of strings.Repeat() function. ◀ */

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go "
	fmt.Print(repeatFunction(str))
}

func repeatFunction(str string) string {
	return strings.Repeat(str, 4)
}
