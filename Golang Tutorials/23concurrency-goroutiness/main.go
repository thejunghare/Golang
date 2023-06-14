package main

import "fmt"
import "time"

func goroutinessFunction() {
	fmt.Println("Hello, World from goroutiness")
}

func main() {
	go goroutinessFunction()
	time.Sleep(1 * time.Second) // This helps to run goroutine
	fmt.Print("Hello World from main function")
}
