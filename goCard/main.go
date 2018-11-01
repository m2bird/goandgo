package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamonds", card()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	
}

func card() string {
	return "Five of Diamonds"
}
