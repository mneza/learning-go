package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp := Employee{"John", "Doe", 55, 6000}
	fmt.Println("First Name:", emp.firstName)
	fmt.Println("Last Name:", emp.lastName)
	fmt.Println("Age:", emp.age)
	fmt.Println(" --- Updated Employee Record --- ")
	updateLastName(&emp)
	fmt.Println("First Name:", emp.firstName)
	fmt.Println("Last Name:", emp.lastName)
	fmt.Println("Age:", emp.age)
}

func updateLastName(employee *Employee) {
	(*employee).lastName = "Test"
}
