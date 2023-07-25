package main


func main() {  
    ch := make(chan int)
    ch <- 5
	// This program will panic as there is no revcive channel
}