package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	fmt.Println("== Structs ==")

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		ContactInfo: ContactInfo{
			email:   "jhalper@dundermifflin.com",
			zipCode: 94000,
		},
	}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
