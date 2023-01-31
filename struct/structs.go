package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Alexey",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 5321,
		},
	}

	alex.newLastName("kwstakis")
	alex.print()
}

func (pointerToPerson *person) newLastName(lastname string) {
	(*pointerToPerson).lastName = lastname
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v\n", pointerToPerson)
}
