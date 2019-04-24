package main

import "fmt"

type Person struct {
	fname string
	lname string
}

type Employee struct {
	fname string
	lname string
	ContactInfo
}

type ContactInfo struct {
	address string
	phone   string
	email   string
}

func main() {
	person := Person{
		fname: "John",
		lname: "Doe",
	}

	fmt.Printf("Person's Name: %v ", person)
	fmt.Println(" --- Employee with ContactInfo ---")
	employee := Employee{
		fname: "John",
		lname: "Doe",
		ContactInfo: ContactInfo{
			address: "test",
			phone:   "123",
			email:   "test@gmail.com",
		},
	}

	fmt.Printf("Employee's Name: %v ", employee)
}
