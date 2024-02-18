package main

func main() {
	cards := new_deck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
