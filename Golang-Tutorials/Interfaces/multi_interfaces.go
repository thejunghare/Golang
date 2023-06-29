package main

import (
	"fmt"
)

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

// inheritance ~ embedding interfaces
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Println("Total Salary:", e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		basicPay:    100,
		pf:          10,
		totalLeaves: 10,
		leavesTaken: 2,
	}

	fmt.Println("Total leaves:", e.CalculateLeavesLeft())
	e.DisplaySalary()

	var employee_operations EmployeeOperations = e
	employee_operations.DisplaySalary()
	fmt.Println(employee_operations.CalculateLeavesLeft())

}