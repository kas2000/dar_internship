package main

import (
	"fmt"
	"kas/card"
)

func main() {
	myDeck := card.NewDeck()
	nums := []string{"Tuz", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	masti := []string{"Piki", "Chervi", "Bubni", "Trefi"}
	for _, num := range nums {
		for _, mast := range masti {
			myDeck.AddCard(card.NewCard(num, mast))
		}
	}

	fmt.Println(myDeck)
	fmt.Println()

	myDeck.ShuffleDeck()
	fmt.Println(myDeck)

	myDeck.RemoveCard("10", "Chervi")
	fmt.Println()

}