/* WAP to find even numbers in a list */

package main

import "fmt"

func main() {
	evenNumbers(10)
}

func evenNumbers(n int) {
	for index := 1; index < n; index++ {
		if index%2 == 0 {
			fmt.Println(index)
		}
	}
}
