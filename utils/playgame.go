package utils

import (
	"BTC/psychic/cards"
	"fmt"
)

var card cards.Card
var Game []cards.Card

// ReadHandAndDeck returns the Hand and the Deck for every new game
func ReadHandAndDeck(gameline []string, lines int) ([]cards.Card, []cards.Card) {
	// file contains game in memory
	var slice1 = Game[:]
	var slice2 = Game[:]
	fmt.Println("number of lines is: ", lines)
	for i := 0; i < lines; i++ {
		for j := 0; j < len(gameline[i]); j = j + 3 { // read character for ch and put it it Card struct
			card.Rank = string(gameline[i][j])
			card.Suit = string(gameline[i][j+1])
			Game = append(Game, card)
		}
	}
	slice1 = Game[0:5]
	slice2 = Game[5:10]
	// we return the Hand and Deck
	return slice1, slice2
}
