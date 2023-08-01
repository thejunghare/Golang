package main

import "fmt"

// creating user defined type Address --> struct
type Address struct { // Address defination
	// defining fields for type address
	name        string
	housenumber int
	street      string
	city        string
	pincode     int
	state       string
}

type User struct {
	userID int
	email  string

	// Now Address is embedded to user
	// so user has access to all fields of Address
	Address

	// Nested struct
	other Address
}

// struct method
func (u User) userDetails(){
	fmt.Println("user name:", u.name)
}

// main function
func main() {
	// declaring a variable of struct type
	var a Address
	fmt.Println(a) // output --> {0,0} (As the variable is only declare and not initalized)

	// Declaring and initalizing (Instance of Address -> address1)
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
	address4 := &Address{
		name:        "John",
		housenumber: 132,
		street:      "New street",
		city:        "la",
		pincode:     400,
		state:       "US",
	}
	fmt.Println((*address4).name)

	// Anonymous struct
	Person := struct {
		firstname string
		lastname  string
	}{
		firstname: "Prasad",
		lastname:  "Junghare",
	}

	fmt.Println(Person)

	
	// Access embedded element
	/* user1 := User {
		Address: Address{
			name: "prasad",
		},
	} */

	user1 := User {	}

	// Short hand to use embedded strut
	user1.state = "USA"
	fmt.Println(user1.state)

	// Modify the values
	user1.Address.name = "john"
	user1.userDetails()

	fmt.Println(user1.name)

	// Access the Nested struct
	nestedUser := User{
		other: Address{
			name: "Prasad",
			housenumber: 31,
		},
	}

	fmt.Println(nestedUser.name)
}
