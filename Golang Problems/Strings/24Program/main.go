/* ▶ Golang code to replace all specified substrings by given substring. ◀ */

package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "Hello world"
	oldString := "world"
	newString := "paddy"
	fmt.Println(replaceString(str, oldString, newString))
}


func replaceString(str string, oldString string, newString string) string {
	return strings.ReplaceAll(str, oldString, newString)
}