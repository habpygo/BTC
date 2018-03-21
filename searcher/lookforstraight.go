package searcher

import (
	"BTC/psychic/cards"
	"fmt"
)

// LookForStraight searches for a straight
func LookForStraight(hand, deck []cards.Card) bool {
	// combine the game
	allCards := append(hand, deck...)
	// Suits are not needed - prepare the Rank slice
	var rankSlice []string
	for _, v := range allCards {
		rankSlice = append(rankSlice, v.Rank)
	}

	fmt.Println("rankSlice is: ", rankSlice)

	return false
}
