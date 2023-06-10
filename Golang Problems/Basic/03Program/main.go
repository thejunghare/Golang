/* GO Program to find area and circumference of circle */

package main

import "fmt"

func main() {
	fmt.Print("circumference of circle is", circumferenceOfCircle(2.4))
}

func circumferenceOfCircle(radius float64) float64 {
	return 2 * 3.14 * radius
}
