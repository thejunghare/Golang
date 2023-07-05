package main

import (
	"fmt"
)


// function to individual digits for 32 -> 2, 3
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		// sending digit to dchnl
		dchnl <- digit
		number /= 10
	}

	// Closing the channel
	close(dchnl)
}

// fuction to calculate square
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)

	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit
	}

	// sending the sum to squareop
	squareop <- 
	
	/* 
		here number recevied will be 2, 3
		hence for 1st number 0 += 2 * 2 -> 4
		similarly for 2nd number 4 += 3 * 3 -> 9 -> 13

		so final output for number = 32 will be 13
	*/
}

// fuction to calculate the cube
func calsCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)

	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit * digit
	}

	cubeop <- sum

	/* 
		here number recevied will be 2, 3
		hence for 1st number 0 += 2 * 2 * 2 -> 8
		similarly for 2nd number 8 += 3 * 3 * 3 -> 27 -> 35

		so final output for number = 32 will be 35
	*/
}

func main() {
	number := 32
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calsCubes(number, cubech)

	// Receive from calcsquares and calccubes in sqrch and cubech
	square := <-sqrch
	cubes := <-cubech

	fmt.Print(square + cubes) // 13+35 = 48
}
