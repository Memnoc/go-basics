package main

func main() {
	cards := new_deck()
	cards.save_to_file("my_cards")
}
