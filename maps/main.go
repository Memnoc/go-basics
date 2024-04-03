package main

import "fmt"

func main() {
	fmt.Println("== Maps ==")

	colors := map[string]string{ // declaring
		"red":   "#ff000",
		"green": "#4b475",
	}

	colors["white"] = "#fffff" // adding

	fmt.Println("Accessing: ", colors["white"]) // accessing

	delete(colors, "white") // deleting key

	for _, color := range colors {
		fmt.Println("Iterating: ", color)
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("printMap(): Hex code for color ", color, "is", hex)
	}
}
