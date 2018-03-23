package searcher

import (
	"BTC/psychic/cards"
	"sort"
	"strings"
)

const NumbeOfCards = 5

var SpadesArray []string
var HeartsArray []string
var DiamondsArray []string
var ClubsArray []string
var Flush bool

func LookForSuits(hand []cards.Card) (string, int, string, int) {
	Spades := 0
	Hearts := 0
	Diamonds := 0
	Clubs := 0

	// Highest and Almost highest possibility Suits are returned together with their lenghts
	var Highest string
	var Number int
	var Almost string
	var Number2 int
	var H string
	var A string
	var LH int
	var LA int

	for i := 0; i < NumbeOfCards; i++ { // hard coded; more than 5 and you cheat :-)
		if strings.Contains(hand[i].Suit, "S") {
			Spades++
			SpadesArray = append(SpadesArray, hand[i].Rank)
		}
		if strings.Contains(hand[i].Suit, "H") {
			Hearts++
			HeartsArray = append(HeartsArray, hand[i].Rank)
		}
		if strings.Contains(hand[i].Suit, "D") {
			Diamonds++
			DiamondsArray = append(DiamondsArray, hand[i].Rank)
		}
		if strings.Contains(hand[i].Suit, "C") {
			Clubs++
			ClubsArray = append(ClubsArray, hand[i].Rank)
		}

	}

	sl := []int{Spades, Hearts, Diamonds, Clubs}
	sort.Ints(sl)

	m := map[string]int{"S": Spades, "H": Hearts, "D": Diamonds, "C": Clubs}

	// sort the map m
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for index, kv := range ss {
		//fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		if index == 0 {
			Highest = kv.Key
			Number = kv.Value
		}
		if index == 1 {
			Almost = kv.Key
			Number2 = kv.Value
		}
	}

	if len(Highest) >= 1 {
		H = Highest
		LH = Number
	}
	if len(Almost) >= 1 {
		A = Almost
		LA = Number2
	}

	return H, LH, A, LA
}

func LookForFlushInDeck(suit string, counter int, deck []cards.Card) bool {
	// index is both a counter and an index
	index := counter
	Spades := 0
	Hearts := 0
	Diamonds := 0
	Clubs := 0

	// check for Suit in deck
	for i := 0; i < NumbeOfCards-index; i++ {
		if strings.Contains(deck[i].Suit, "S") {
			Spades++
		}
		if strings.Contains(deck[i].Suit, "H") {
			Hearts++
		}
		if strings.Contains(deck[i].Suit, "D") {
			Diamonds++
		}
		if strings.Contains(deck[i].Suit, "C") {
			Clubs++
		}

		switch suit {
		case "S":
			SpadesArray = append(SpadesArray, deck[i].Rank)
		case "H":
			HeartsArray = append(HeartsArray, deck[i].Rank)
		case "D":
			DiamondsArray = append(DiamondsArray, deck[i].Rank)
		case "C":
			ClubsArray = append(ClubsArray, deck[i].Rank)
		default:
		}
	}
	switch suit {
	case "S":
		counter += Spades
	case "H":
		counter += Hearts
	case "D":
		counter += Diamonds
	case "C":
		counter += Clubs
	}
	if counter == 5 { // we have 5 equal Suits at least
		Flush = true // so at least there is a Flush
		if suit == "S" {
			StraightFlush := IsFlushStraight(SpadesArray)
			if StraightFlush {
				Flush = false
				return true
			}
		}
		if suit == "H" {
			StraightFlush := IsFlushStraight(HeartsArray)
			if StraightFlush {
				Flush = false
				return true
			}
		}
		if suit == "D" {
			StraightFlush := IsFlushStraight(DiamondsArray)
			if StraightFlush {
				Flush = false
				return true
			}
		}
		if suit == "C" {
			StraightFlush := IsFlushStraight(ClubsArray)
			if StraightFlush {
				Flush = false
				return true
			}
		}
		return false
	}
	// clean out Suit arrays for next game
	SpadesArray = SpadesArray[:0]
	HeartsArray = HeartsArray[:0]
	DiamondsArray = DiamondsArray[:0]
	ClubsArray = ClubsArray[:0]

	return false
}

// IsFlushStraight looks if we are dealing with a straight-flush or normal flush
func IsFlushStraight(rankarray []string) bool {
	// compare rankarray with cards
	RankedMap := make(map[string]int)
	// map to int values to sort according to card game
	for i := 0; i < NoOfCards; i++ {
		RankedMap[rankarray[i]] = cards.SetRank(rankarray[i])
	}

	type pokerhand struct {
		Key   string
		Value int
	}

	var structslice []pokerhand // slice of struct kv
	for rank, suit := range RankedMap {
		structslice = append(structslice, pokerhand{rank, suit})
	}

	sort.Slice(structslice, func(i, j int) bool {
		return structslice[i].Value > structslice[j].Value
	})

	var stringSlice []string
	for _, kv := range structslice {
		stringSlice = append(stringSlice, kv.Key)
	}
	stringSet := strings.Join(stringSlice, "")

	// now see whether it compares as a subslice of RankSet
	if strings.Contains(cards.StringSet, stringSet) {
		return true
	}

	return false
}
