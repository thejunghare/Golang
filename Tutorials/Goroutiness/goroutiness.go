package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	go world()
	time.Sleep(3 * time.Second)
	/*
		In this case using the time package the main function is put to sleep for 1 second and go hello()
		has enough time to execute before the main goroutine
		terminates.
	*/
	fmt.Println("Main function")
}

func hello() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func world() {
	for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
