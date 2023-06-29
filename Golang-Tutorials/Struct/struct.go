package main

import "fmt"

// creating user defined type Address --> struct
type Address struct {
	// defining fields for type address
	name        string
	housenumber int
	street      string
	city        string
	pincode     int
	state       string
}

// main function
func main() {
	// declaring a variable of struct type
	var a Address
	fmt.Println(a) // output --> {0,0} (As the variable is only declare and not initalized)

	// Declaring and initalizing
	address1 := Address{
		"Jule",
		132,
		"Time Square",
		"New York",
		400,
		"US",
	}
	fmt.Println("Fields of address1:", address1) // output --> {Jule 132 Time Square New York 400 US}

	// Name fields while initalizing
	address2 := Address{
		name:        "Jule",
		housenumber: 132,
		street:      "Time Square",
		city:        "New York",
		pincode:     400,
		state:       "US",
	}

	fmt.Println("Fields with value:", address2) // output --> {Jule 132 Time Square New York 400 US}

	// uninitalized fields will be set to corresponding zero-value
	address3 := Address{
		name: "Jule",
	}
	fmt.Println(address3) // output --> {Jule, 0}

	// Access the fields of type address
	fmt.Println("City name is: ", address1.city) // output --> City name is: New York

	// using pointers with srtuct
	address4 := &Address {
		name:        "John",
		housenumber: 132,
		street:      "New street",
		city:        "la",
		pincode:     400,
		state:       "US",
	}
	fmt.Println((*address4).name)
}
