/* WAP to swap values */

package main

import "fmt"

func main() {
	x := 3
	y := 4

	fmt.Println(swapValuesWithOutThirdVariable(x, y))

	fmt.Println(usingDestrucuringAssignment(x, y))

	fmt.Print(swapValuesWithOutThirdVariable(x, y))
}

func swapValuesWithThirdVariable(x int, y int) (int, int) {
	temp := 0
	temp = x
	x = y
	y = temp
	return x, y
}

func swapValuesWithOutThirdVariable(x int, y int) (int, int) {
	x = x + y
	y = x - y
	x = x - y
	return x, y
}

func usingDestrucuringAssignment(x int, y int) (int, int) {
	x, y = y, x
	return x, y
}
