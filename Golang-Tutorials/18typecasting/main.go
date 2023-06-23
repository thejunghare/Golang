/*
	Type casting: Converting the value of one data type into another datatype

	1. Explicit type casting: manually type casting.
	2. Implicit type casting: automatic type casting.

	"GO DOESN’T SUPPORT IMPLICIT TYPE CASTING"
*/

package main

import "fmt"

func main() {
	// fmt.Println(explicitTypeCasting())
	explicitTypeCasting()

	fmt.Println("sum is ", implicitTypeCasting())
}

/* Explicit type casting */

func explicitTypeCasting() {
	floatingValue := 9.8
	fmt.Printf("Type of floatingValue before type casting is: %T", floatingValue)

	fmt.Println("")

	integerValue := int(floatingValue)

	fmt.Printf("Type of floatingValue after type casting: %T", integerValue)

	fmt.Println("")

}

func implicitTypeCasting() float64 {
	/*
		initialize integer variable to a floating-point number
		var number int = 4.34

		fmt.Printf("Number is %T", number)

		This will through error as implict convert isn't support by golang
	*/

	var integerValue int = 12

	var floatingValue float64 = 3.14

	var sum float64 = (float64)(integerValue) + (floatingValue)

	return sum
}
