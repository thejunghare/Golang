package main

import "fmt"

func channelFunction(ch chan int) {
	fmt.Println(10 + <-ch)
}

func main(){
	fmt.Println("Start main function")

	// Creating an channel
	ch := make(chan int)

	go channelFunction(ch)
	ch <- 40
}
