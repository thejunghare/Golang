/* WAP to find the given string is Palindrome or not */

package main

import "fmt"

func main() {
	str := "BOB"
	fmt.Print(isPalindrome(str))
}

func isPalindrome(str string) bool {
	var result string
	for _, v := range str {
		result = string(v) + result
	}

	if str == result {
		return true
	}

	return false
}
