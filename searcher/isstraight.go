package searcher

import (
	"BTC/psychic/cards"
	"fmt"
	"strings"
)

var straight bool = false
var straightslice []string
var value1 []string
var value2 []string
var value3 []string

func IsSimpleStraight(hand, deck []cards.Card) bool {
	RankedStraight := make(map[string]int)

	// sort the rawhand
	//var intkeeper []int
	for i := 0; i < NoOfCards; i++ {
		RankedStraight[hand[i].Rank] = cards.SetStraightRank(hand[i].Rank)
		//intkeeper = append(intkeeper, RankedStraight[hand[i].Rank])
	}

	newPairlist := SortMap(RankedStraight)
	fmt.Println("newPairlist is: ", newPairlist)
	for k, v := range newPairlist {
		fmt.Println(" k, v: ", k, v)
		value1 = append(value1, v.Key)
	}

	//fmt.Println("houder is:", intkeeper)
	// load the slices and check
	//for i := 0; i < NoOfCards; i++ {

	//value1 = append(value1, hand[i].Rank)

	//fmt.Println("value1:", value1)

	for j := 0; j < NoOfCards-i; j++ {
		value2 = append(value2, deck[j].Rank)
		//	fmt.Println("value2:", value2)

		// sort value1

		straightslice = append(value1, value2...)
		//fmt.Println("straightslice: ", straightslice, "len is: ", len(straightslice))
		sortedSlice := MapSorter(straightslice, len(straightslice))

		fmt.Println("sortedSlice: ", sortedSlice)

		value2 = value2[:0]
		value1 = value1[:0]

		if len(straightslice) == 5 {

			returnSet := MapSorter(straightslice, len(straightslice))
			if len(returnSet) == 5 && strings.Contains(cards.StringSet, returnSet) || strings.Contains(cards.StraightSet, returnSet) {
				//		fmt.Println("we found it")
				straight = true

			}
		}
		straightslice = straightslice[:0]
		//value1 = value1[:i+1]
	}
	//value1 = value1[:0]

	for k := range RankedStraight {
		delete(RankedStraight, k)
	}
	straight = false

	return straight
}
