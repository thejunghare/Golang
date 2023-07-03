/* Chef and Chefina are playing with dice. In one turn, both of them roll their dice at once.

They consider a turn to be good if the sum of the numbers on their dice is greater than
6
6.
Given that in a particular turn Chef and Chefina got
�
X and
�
Y on their respective dice, find whether the turn was good. */

package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		x, y := takeinput()

		if x+y > 6 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func takeinput() (int, int) {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	return x, y
}
