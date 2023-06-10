/* GO Program to Check Whether a Number is Palindrome or Not */

package main

import "fmt"

func isPalindrome(number int) bool {
	reverse := 0

	for number > 0 {
		remainder := number % 10
		reverse = (reverse * 10) + remainder
		number /= 10
	}

	return reverse == number
}

func main() {
	fmt.Println(isPalindrome(1031))
	fmt.Print(isPalindrome(88))
}
