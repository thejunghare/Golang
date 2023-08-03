package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else if val == ')' || val == '}' || val == ']' {
			if len(stack)==0||stack[len(stack)-1] != brackets[val]{
				return false
			}
			stack= stack[:len(stack)-1]
		} 
	}


	return len(stack)==0
}

func main() {
	fmt.Println(isValid("("))
}
