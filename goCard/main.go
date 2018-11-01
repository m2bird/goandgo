package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamonds", card()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	
	for i, card := range cards{
		fmt.Println(i,card)
	}
}

func card() string {
	return "Five of Diamonds"
}
