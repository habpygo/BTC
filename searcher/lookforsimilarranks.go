package searcher

import (
	"BTC/psychic/cards"
	"fmt"
)

func LookForFour(rank string, noOfRanks int, deck []cards.Card) (bool, bool, bool) {

	// look also for three of a kind or one pair
	fmt.Println("")
	return false, true, false // or false,false,true or true,false,false etc.
}

func LookForFullHouse(rank1 string, rank2 string, nor1 int, nor2 int, deck []cards.Card) bool {

	// look also for two-pair
	fmt.Println("")
	return false
}

func LookForStraight(rank string) bool {

	// look also for highest card
	fmt.Println("")
	return false
}
