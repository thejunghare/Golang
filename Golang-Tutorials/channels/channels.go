package main

import "fmt"

func main() {
	var mychannel chan int
	fmt.Println(mychannel)
	fmt.Printf("Type of channel %T", mychannel)

	// Creating channel using make() function
	mychannel1 := make(chan int)
	fmt.Println(mychannel1)
	fmt.Printf("Type of channel %T", mychannel1)
}
