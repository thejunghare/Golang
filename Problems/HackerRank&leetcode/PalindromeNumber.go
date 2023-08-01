package main

import "fmt"

func isPalindrome(x int) bool {
	reverse := 0
	temp := x

	for temp > 0 {
		remainder := temp % 10
		reverse = (reverse * 10) + remainder
		temp /= 10
	}

	if x == reverse {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome(121))
}
