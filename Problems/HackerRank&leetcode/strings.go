package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ModifyString(str string) string {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	var builder strings.Builder

	for _, r := range result {
		if !unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	removeIntgeres := builder.String()
	trimspace := strings.Trim(removeIntgeres, " ")
	return trimspace
}

func main() {
	str := " de75s1rev2er "
	fmt.Print(ModifyString(str))
}
