/* ▶ Golang program to get the length of the specified string ◀ */
package main

import "fmt"

func main() {
	fmt.Print(findLength("str"))
}

func findLength(str string) int {
	return len(str)
}
