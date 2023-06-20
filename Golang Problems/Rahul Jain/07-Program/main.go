/* WAP to reverse a string */

package main

import "fmt"

func main() {
	str := "Paddy"
	fmt.Print(reverse_string(str))
}

func reverse_string(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}

/* func usingForLoop(str []string) {
	right := len(str) - 1
	for index := 0; index < right; index++ {
		temp := str[index]
		str[index] = str[(len(str) - 1)]
		str(right) = temp
	}
} */
