package main

import (
	// "fmt"
	"time"
)

func Task1() {
	x := 5
	x++
	println(x)
}

func Task2() {
	y := 6
	y--
	println(y)
}

func main() {
	go Task1()
	go Task2()
	time.Sleep(3 * time.Second)
	/*
		In this case using the time package the main function is put to sleep for 1 second and go hello()
		has enough time to execute before the main goroutine
		terminates.
	*/
	println("Main function")
}
