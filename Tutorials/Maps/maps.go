package main

import (
	"fmt"
)

// Create Map
func createMap() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)
}

// Add items to Map
func addItems() {
	employeeSalary := make(map[string]int)

	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000

	fmt.Println("employeeSalary map contents:", employeeSalary)
}

// Add items while declaration
func addWhileDeclaration() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}

	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)
}

// retriveing value from keys
func retrievingValueFromKey() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}

	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}

// Check if keys exists
func isPresent() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000, "mike": 9000,
	}

	checkKey := "steve"
	value, ok := employeeSalary[checkKey]
	if ok == true {
		fmt.Println("Salary of", checkKey, "is", value)
		return
	}

	fmt.Println(checkKey, "not found")
}

// Delete map items
func deleteMapItems() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000, "mike": 9000,
	}

	fmt.Println("Length is", len(employeeSalary))
	delete(employeeSalary, "steve")
	fmt.Println("employeeSalary map contents:", employeeSalary)

	fmt.Println("Length is", len(employeeSalary))

}

// Print all elements
func iterateOverAllElements() {
	empolyeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}

	for key, value := range empolyeeSalary {
		fmt.Printf("%s = %d\n", key, value)
	}
}

// Struct with maps
type employee struct {
	salary  int
	country string
}

func main() {
	
	/* bike := map[string]string{"Brand": "Honda", "Model": "cb shine"}
	fmt.Println(bike) */


	my_map := make(map[string]string) // my_map is empty ->  make function is the correct way to create an empty map
	my_map["brand"] = "Honda"         // my_map is no longer empty

	/*
		Create empty map without using make function

		var map_name map[keyType]ValueType

		Note: Create an empty map  fucntion without make function will cause a runtime panic
			panic -> something went unexpectedly wrong
	*/

	createMap()
	addItems()
	addWhileDeclaration()
	retrievingValueFromKey()
	isPresent()
	deleteMapItems()
	iterateOverAllElements()

	emp1 := employee{
		salary:  12000,
		country: "usa",
	}

	employeeInfo := map[string]employee{
		"Steve": emp1,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary: $%d country: %s\n", name, info.salary, info.country)
	}
}
