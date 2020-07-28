package card

import "math/rand"
import "time"

type deck struct {
	cards []card
}

//NewDeck cosntructor
func NewDeck() deck {
	return deck {
		cards: make([]card, 0),
	}
}

//helper function for AddCard to check for duplicates in a deck
func hasDuplicate(d *deck, c card) bool {
	for _, v := range d.cards {
		if (v == c) {
			panic("Such element already exists in a deck")
		}
	}

	return false;
}

func outFromCapacity(d *deck) bool {
	if (len(d.cards) >= 52) {
		panic("Your deck will exceed the limit of 52 cards")
	}

	return false;
}

//AddCard adds a card to the deck
func (d *deck) AddCard(c card) {
	if (hasDuplicate(d, c) == false && outFromCapacity(d) == false) {
		d.cards = append(d.cards, c)
	}
}

//Removes a card from the deck
func (d *deck) RemoveCard(num string, lear string) {
	cardHasBeenFound := false
	for	i, value := range d.cards {
		// fmt.Println(value.Num())
		if (value.Num() == num && value.Lear() == lear) {
			copy(d.cards[i:], d.cards[i+1:])
			d.cards[len(d.cards)-1] = card{num: "0", lear: "not defined"}
			d.cards = d.cards[:len(d.cards)-1]
			cardHasBeenFound = true
			break;
		}
	}
	if (!cardHasBeenFound) {
		panic("Such element was not found in a deck")
	}
}

//Shuffle is a function which shuffles the deck
func (d *deck) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

//Cards is a getter for a copy of cards
func (d deck) CardsCopy() []card {
	cardsCopy :=  make([]card, len(d.cards))
	copy(cardsCopy, d.cards)
	return cardsCopy
}

//Sets the new cards to the deck
func (d *deck) SetCards(newCards []card) {
	d.cards = newCards
}

//Tail returns the last element from the deck
func (d deck) Tail() card {
	return d.cards[len(d.cards) - 1]
}

//Head return the first element from the deck
func (d deck) Head() card {
	return d.cards[0]
}