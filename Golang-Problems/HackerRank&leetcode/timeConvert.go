package main

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {
	format12 := "03:04:05PM"
	format24 := "15:04:05"

	// Parse the input time string using the 12-hour format
	t, err := time.Parse(format12, s)
	if err != nil {
		panic(err)
	}

	// Format the time to the desired 24-hour format
	return t.Format(format24)
}

func main() {
	s := "04:04:05PM"
	fmt.Print(timeConversion(s))
}
