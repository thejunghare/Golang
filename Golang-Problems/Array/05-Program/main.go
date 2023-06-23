/* WAP to print second largest element in array without sorting the array */

package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3}

	largest := arr[0]

	for index := 0; index < len(arr); index++ {
		if arr[index] > largest {
			largest = arr[index]
		}
	}

	fmt.Print("Largest element:", largest)
}

