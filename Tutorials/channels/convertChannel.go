package main

import "fmt"

// This function converts bidirectional channel to send only channel(unidirectional channel)
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	// creating a bidirectional channel.
	chnl := make(chan int)

	// The senddata function converts this chnl to send only channel.
	go sendData(chnl)

	// channel chnl is bidirectional in main goroutine so this program will print 10 as the output.
	fmt.Print(<-chnl)
}
