package searcher

import (
	"BTC/psychic/cards"
	"strings"
)

// FindTotalDeckRanks finds the number of Ranks in the Deck depending on the number of Ranks found in Hand
func FindTotalDeckRanks(suit string, handtotal int, deck []cards.Card) int {
	totalInDeck := 0
	for i := 0; i < handtotal-1; i++ { // we cater for only one found in hand
		totalInDeck += strings.Count(deck[i].Rank, suit)
	}
	return totalInDeck
}
