package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main(){
	p := &Point{
		X:1,
		Y:2,
	}

	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Print(p)
}