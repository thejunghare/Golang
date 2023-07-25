/* The task is to find their comparison points by comparing a[0] with b[0], a[1] with b[1], and a[2] with b[2].

If a[i] > b[i], then Alice is awarded 1 point.
If a[i] < b[i], then Bob is awarded 1 point.
If a[i] = b[i], then neither person receives a point.
Comparison points is the total points a person earned.

Given a and b, determine their respective comparison points. */

package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	alicepoints := int32(0)
	bobpoints := int32(0)

	var c = []int32{}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alicepoints++
		} else if a[i] < b[i] {
			bobpoints++
		}
	}

	c = append(c, alicepoints, bobpoints)
	return c
}

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}

	fmt.Print(compareTriplets(a, b))

}
