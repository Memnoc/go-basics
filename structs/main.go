package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func main() {
	fmt.Println("== Structs ==")

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: ContactInfo{
			email:   "jhalper@dundermifflin.com",
			zipCode: 94000,
		},
	}

	fmt.Println("My first name is: ", jim.firstName)
	fmt.Println("My last name is: ", jim.lastName)
	fmt.Printf("Person info: %+v\n", jim)
}
