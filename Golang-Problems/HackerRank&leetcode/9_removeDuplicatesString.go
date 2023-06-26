package main

import "fmt"

func removeDuplicates(str []string) []string {
	maps := make(map[string]bool)
	var slices []string

	for _, val := range str {
		if_, ok := maps[val]; !ok {
			maps[val] = true
			slices = append(slices, val)
		}
	}
	return slices
}

func main() {
	str := []string{"apple", "banana", "apple", "orange", "banana", "kiwi"}
	fmt.Print(removeDuplicates(str))
}