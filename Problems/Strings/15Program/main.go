/* ▶ Printing the ASCII value of characters of the string in Golang. ◀ */

package main

import "fmt"

func main() {
	str := "hello"
	printAsciiValues(str)
}

func printAsciiValues(str string) {
	for index := 0; index < len(str); index++ {
		fmt.Printf("%c %d\n", str[index], str[index])
	}
}
