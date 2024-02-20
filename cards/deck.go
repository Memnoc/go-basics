package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

// A deck type which is a slice of string
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

/*
* NOTE: utils to read from a file, just for fun
 */
func (d deck) read_from_file() error {
	content, err := os.ReadFile("../cards/my_cards")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("File contents: %s", content)
	return err
}

func new_deck_from_file(filename string) deck {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("No such file or directory", err)
		os.Exit(1)
	}

	s := strings.Split(string(content), ",")
	return deck(s)
}
