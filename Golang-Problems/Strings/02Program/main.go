/*
	Check a string contains a specified substring
	str := "Hello" subString := "Wor"
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	subString := "Wor"

	fmt.Print(checkIfContains(str, subString))

}

func checkIfContains(str string, subString string) bool {
	return strings.Contains(str, subString)
}
