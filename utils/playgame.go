package utils

import (
	"BTC/psychic/cards"
	"fmt"
)

var card cards.Card
var Game []cards.Card
var Slice1 = Game[:]
var Slice2 = Game[:]

// ReadHandAndDeck returns the Hand and the Deck for every new game
func ReadHandAndDeck(gameline []string, lines int) ([]cards.Card, []cards.Card) {
	// file contains game in memory

	fmt.Println("number of lines is: ", lines)

	for i := 0; i < lines; i++ {
		for j := 0; j < len(gameline[i]); j = j + 3 { // read character for ch and put it it Card struct
			card.Rank = string(gameline[i][j])
			card.Suit = string(gameline[i][j+1])
			Game = append(Game, card)
		}
	}
	Slice1 = Game[0:5]
	Slice2 = Game[5:]
	// we return the Hand and Deck
	return Slice1, Slice2
}
