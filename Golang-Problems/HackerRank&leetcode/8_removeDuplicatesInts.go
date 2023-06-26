package main

func removeDuplicates(nums []int) []int {
	maps := map[int]bool{}
	empty := []int{}

	for element := range nums {
		if maps[nums[element]] != true {
			maps[nums[element]] = true
			empty = append(empty, nums[element])
		}
	}

	return empty
}

func main() {
	nums := []int{"apple", "banana", "apple", "orange", "banana", "kiwi"}
	fmt.Println(removeDuplicates(nums))
}
