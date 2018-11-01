package main

import "fmt"

func main() {

	cards := deck{"Ace of Diamonds", card()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	
	cards.print()
}

func card() string {
	return "Five of Diamonds"
}
