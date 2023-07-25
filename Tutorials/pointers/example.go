package main

import "fmt"

func main() {
	x := 0
	y := 0

	fmt.Printf("before:\ta=%d\tb=%d\n", x, y)
	updated(x, &y)
	fmt.Printf("after:\ta=%d\tb=%d\n", x, y)

	// Get Address -> we need to put ampersand(&) in front of variable to get address
	num := 10
	fmt.Println("address of variable num", &num)

	// Get value -> use the * opertor
	num2 := &num
	fmt.Println(*num2)
}

func updated(a int, b *int) {
	a = 10
	*b = 10
	fmt.Printf("inside:\ta=%d\tb=%d\n", a, *b)
}
