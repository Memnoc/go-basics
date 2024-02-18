package main

func main() {
	cards := deck{
		new_card(), "Ace of Diamonds",
	}
	cards = append(cards, "Ace of Hearts")

	cards.print()
}

func new_card() string {
	return "Five of Spades"
}
