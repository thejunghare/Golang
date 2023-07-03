/* Given a square matrix, calculate the absolute difference between the sums of its diagonals. */

package main

import "fmt"

func diagonalDifference(a [][]int32) int32 {
	primarydiagonal := int32(0)
	secondarydiagonal := int32(0)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if i == j {
				primarydiagonal += a[i][j]
			}

			if (i + j) == (len(a) - 1) {
				secondarydiagonal += a[i][j]
			}
		}
	}
	return secondarydiagonal - primarydiagonal
}

func main() {
	a := [][]int32{
		{-1, 1, -7, -8},
		{-10, -8, -5, -2},
		{0, 9, 7, -1},
		{4, 4, -2, 1},
	}

	fmt.Print(diagonalDifference(a))

}
