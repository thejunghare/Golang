package main

import "fmt"

func channelFunction(ch chan int) {
	fmt.Println(10 + <-ch)
}

func main() {
	fmt.Println("Start main function")

	// Creating an channel
	ch := make(chan int) // If you use this syntax to check the len and cap of channel it will always return '0' because the channel operation ch <- 40 will block the channel

	// ch := make(chan int, 1) // use this buffered channel an buffered channel allows sending and receiving values without immediately blocking as long as the channel buffer is not full.
	
	go channelFunction(ch)
	ch <- 40
	fmt.Println("Length of ch channel is ", len(ch))
	fmt.Println("Capacity of ch channel is ", cap(ch))

	// closing an channel
	close(ch)
}
