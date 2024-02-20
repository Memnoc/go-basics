package main

func main() {
	cards := new_deck_from_file("../cards/my_cards")
	cards.shuffle()
	cards.print()
	// cards := new_deck()
	// cards.save_to_file("my_cards")
	// cards.read_from_file()
}
