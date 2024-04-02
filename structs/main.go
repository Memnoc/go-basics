package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("== Structs ==")

	jim := person{"Jim", "Halpert"}

	fmt.Println("My first name is: ", jim.firstName)
	fmt.Println("My last name is: ", jim.lastName)
}
