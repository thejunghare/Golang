/* ▶ Golang program to match two same strings in different cases. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "INDIA"
	anotherString := "india"
	fmt.Println(sameStringInDifferentCase(str, anotherString))
}

func sameStringInDifferentCase(str string, anotherString string) bool {
	return strings.EqualFold(str, anotherString)
}
