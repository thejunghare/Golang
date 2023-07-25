/* GO Program to Find the Largest Number Among Three Numbers */
package main

import "fmt"

func main() {
	first := 1223
	second := 131
	third := 14
	largestNumber(first, second, third)
}

func largestNumber(first int, second int, third int) {
	if first >= second && first >= third {
		fmt.Print("First is largest")
	}

	if second >= first && second >= third {
		fmt.Print("Second is largest")
	}

	if third >= first && third >= second {
		fmt.Print("Third is largest")
	}
}
