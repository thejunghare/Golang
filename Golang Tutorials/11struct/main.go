package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	job string
}

func main (){
	/* 
		struct -> structures are variables used to store different types in one variable

		type struct_name struct {
			member1 datatype;
			member2 datatype;
			member3 datatype;
		}
	*/

	/* Access struct members */
	var person1 Person

	person1.name = "Prasad"
	person1.age = 22
	person1.job = "Golang Dev"

	fmt.Println("Name:",person1.name)

	/* Access through function */
	printPerson(person1)
}

func printPerson(person1 Person){
	fmt.Println("Age printed through function:", person1.age)
}