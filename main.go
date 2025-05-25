package main

import "fmt"

func main() {
	cards := newDeck()
	cards.showCards()
	hand, left := deal(cards, 4)
	fmt.Println(hand, left)

	cards.saveFile("cardfile")
}

