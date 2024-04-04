package main

import "fmt"

type bot interface {
	getGreeting() string
}

type english struct {
	greeting string
}

type spanish struct {
	greeting string
}

func (e english) getGreeting() string {
	return e.greeting
}

func (s spanish) getGreeting() string {
	return s.greeting
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	fmt.Println("== Interfaces ==")

	eng := english{greeting: "Hello"}
	spa := spanish{greeting: "Hola"}

	printGreeting(eng)
	printGreeting(spa)
}
