/* ▶ Golang code to get characters from a string using the index. ◀ */
package main

import "fmt"

func main() {
	accessCharacter("Hello")
}

func accessCharacter(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
}
