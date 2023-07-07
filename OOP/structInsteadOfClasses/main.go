package main

import "/oop/employee"

func main()  {
	obj := employee.Employee {
		FirstName: "John",
		LastName: "Cena",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	obj.LeavesRemaining()

	/* --- 	New() instead of using constructors	---  */

	/* New() function of constructors */
	// var e employee.Employee // We have create a zero value of Employee
	// e.LeavesRemaining()

	e := employee.New("Randy", "ortan", 30, 20)
	e.LeavesRemaining()
}