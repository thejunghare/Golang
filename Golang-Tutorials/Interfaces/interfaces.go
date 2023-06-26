package main

import (
	"fmt"
	"math"
)

/*
	Interfaces: 1. Custom type used to specfiy a set of one or more method signatures
				2. interfaces are abstract.
				3. You are not allowed to create an instance of the interface.
				4. Its necessary to implement all the methods declared in interface.

	syntax: type interface_name interface {
			method_name() string
	}
*/

/* Interface for shape */
type shape interface {
	area() float64
	circumf() float64
}

/* struct square */
type square struct {
	length float64
}

/* Methods for square */
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

/* struct circle */
type circle struct {
	radius float64
}

/* Methods for circle */
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape){
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())
}

/* Main method */
func main() {
	shapes := []shape {
		square{length: 15.2},
		circle{radius: 12.3},
	}

	for _,v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
