/* ▶ Golang code to split the string using a specified delimiter ◀ */

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello-world"
	splitString(str)
}

func splitString(str string) {
	var arr []string = strings.Split(str, "-")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
