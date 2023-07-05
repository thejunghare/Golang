package main

import (
	"fmt"
	// "time"
)

// producer function write 0 to 9 to the channel
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		// sending values to chn form 0 to 9
		chnl <- i
	}

	// closing the channel
	close(chnl)
}

// using range keyword to receive the data
func receive(chnl chan int) {
	for v := range chnl {
		fmt.Println("Received through range ", v)
	}

}

func main() {
	// create an channel
	ch := make(chan int)

	// calling the producer function
	go producer(ch)

	// receiving through receive function
	// go receive(ch)

	// Below is an infinite for loop to check if the channel is closed or not
	for {
		// v is receiving the data
		v, ok := <-ch
		if ok == false {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received", v, ok)
	}
}
