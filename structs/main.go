package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	fmt.Println("== Structs ==")

	jim := Person{
		firstName: "Jim",
		lastName:  "Halpert",
		ContactInfo: ContactInfo{
			email:   "jhalpert@dundermifflin.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimbo")
	jim.print()
}

func (p *Person) updateName(name string) {
	p.firstName = name
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
