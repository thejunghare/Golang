package main

import (
	"fmt"
)

func main() {
	arr := []int{0,1,2,2,3,0,4,2}
	val := 2

	var result []int

	for _, key := range arr {
		if key != val {
			result = append(result, key)
		}
	}

	fmt.Println(len(result))
}
