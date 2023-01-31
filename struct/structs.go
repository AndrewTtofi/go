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
	fmt.Printf("%+v\n", alex)
}
