package main

import "fmt"

// struct - employee
type Employee struct {
	name     string
	salary   int
	currency string
}

/* Method - displaySalary() */
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

/* Main function */
func main() {
	employee1 := Employee{
		name:     "John Cena",
		salary:   5000,
		currency: "$",
	}

	/* Calling method */
	employee1.displaySalary()
}
