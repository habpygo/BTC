package searcher

import (
	"BTC/psychic/cards"
	"strconv"
	"strings"
)

var straight bool = false
var straightslice []string
var valuedeck []string

// IsStraight looks if a Straight can be formed with the hand and n cards from the deck. It returns a bool.
func IsStraight(rawhand, deck []cards.Card) bool {

	RankStraight := make(map[string]int)
	hand := make([]cards.Card, len(rawhand))

	// first we map the Ranks to values according to the specs in
	// package cards in such a way that we can sort them correctly
	for i := 0; i < NoOfCards; i++ {
		RankStraight[rawhand[i].Rank] = cards.SetStraightRank(rawhand[i].Rank)
	}

	// sort the map before we pass it on to hand
	newPairlist := SortMap(RankStraight)

	for k, v := range newPairlist {
		hand[k].Rank = v.Key
		hand[k].Suit = strconv.Itoa(v.Value)
	}

	valueslice := []string{}
	straightslice := []string{}
	valuedeck := []string{}

	for i := 0; i < NoOfCards-1; i++ {
		valueslice = AppendString(valueslice, hand[i].Rank)

		// fill the deck and add the rest from the deck
		for j := 0; j < NoOfCards-i-1; j++ {
			valuedeck = AppendString(valuedeck, deck[j].Rank)
		}
		straightslice = AppendString(valueslice, valuedeck...)
		// analyze straightslize and restard
		if len(straightslice) == 5 {
			returnSet := MapSorter(straightslice, len(straightslice))
			if strings.Contains(cards.StringSet, returnSet) || strings.Contains(cards.StraightSet, returnSet) { // check for A==14 or A==1
				straight = true
				continue
			}
			straight = false
		}
		valuedeck = valuedeck[:0]
	}

	// delete map
	for k := range RankStraight {
		delete(RankStraight, k)
	}

	return straight
}

// AppendString appends data to the end of a slice.
// This function appends string elements to a slice of bytes,
// growing the slice if necessary, and returns the updated slice value
func AppendString(slice []string, data ...string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
