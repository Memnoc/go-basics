package main

import (
	"fmt"
	"io/fs"
	"log"
	"strings"

	// A deck type which is a slice of string
	"os"
)

type deck []string

func new_deck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) save_to_file(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), fs.FileMode(0666))
}

func (d deck) save_to_file_two(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// func test() {
// 	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
