/* 
	Given a string, return its reverse.
*/

package main

import "fmt"

func Reverse(word string) string {
	var res string

	for _, val := range word {
		res = string(val) + res
	}
	return res
}

func ClassicReverse(word string) string {
	/* NOTE: THIS FUNC WILL FAIL FOR CHAR THAT TAKES MORE THAN 1 BYTE SPACE */
	var res string

	for i :=  len(word) -1 ; i >= 0; i-- {
		res = res + string(word[i])
	}

	/* 
	This will also work
	for i := 0; i < len(word); i++ {
		res = string(word[i]) + res
	} */

	return res
}

func main() {
	word := "Hello"
	fmt.Println(Reverse(word))
	fmt.Print(ClassicReverse(word))
}
