package main

import "fmt"

func main() {

	/* Single case */
	day := 4

	switch day {
	case 1:
		fmt.Print("Monday")
	case 2:
		fmt.Print("Tuesday")
	case 3:
		fmt.Print("Wednesday")
	case 4:
		fmt.Print("Thursday")
	case 5:
		fmt.Print("Friday")
	case 6:
		fmt.Print("Saturday")
	case 7:
		fmt.Print("Sunday")
	default:
		fmt.Print("Invalid day of number")
	}

	fmt.Println("")

	/* Multi-case */
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of number")
	}
}
