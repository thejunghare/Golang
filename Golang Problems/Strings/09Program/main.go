/* ▶ Golang code to demonstrate the example of strings.Replace() function. ◀ */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(replaceFunction("india"))
}

func replaceFunction(str string) string {
	/*
		strings.Replacce(string, "old string", "new string", -1//1)
		-> 1 will replace string at first occurences
		-> -1 will replace string at all occurences
	*/
	return strings.Replace(str, "i", "I", 1)
}
