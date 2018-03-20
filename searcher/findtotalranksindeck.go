package searcher

import (
	"BTC/psychic/cards"
	"strings"
)

var TotalInDeck int = 0

// FindTotalDeckRanks finds the number of Ranks in the Deck depending on the number of Ranks found in Hand
func FindTotalDeckRanks(suit string, handtotal int, deck []cards.Card) int {
	TotalInDeck := 0
	for i := 0; i+1 <= 5-handtotal; i++ { // we cater for only one found in hand
		TotalInDeck += strings.Count(deck[i].Rank, suit)
	}
	return TotalInDeck
}

// IF FIRST SECOND THIRD FOURTH AND FITHT ARE 1 -> STRAIGHT
