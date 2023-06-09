/* ▶ Golang code to find the index of last occurrence of specified character in specified string. ◀ */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ABC PQR LMN PQR"
	byteValue := byte('Q')

	fmt.Println(lastOccurrence(str, byteValue))
}

func lastOccurrence(str string, byteValue byte) int {
	return strings.LastIndexByte(str, byteValue)
}
