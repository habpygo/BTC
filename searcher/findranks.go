package searcher

import (
	"BTC/psychic/cards"
	"sort"
	"strings"
)

const NoOfCards = 5

var deckExchange int

// FindRanks will give the maximum result of the hand and deck input given the number of similars found in hand
// this will dictate the iterations over the deck
func FindRanks(hand, deck []cards.Card) int {

	HandAce := 0
	HandKing := 0
	HandQueen := 0
	HandJack := 0
	HandT := 0
	HandNine := 0
	HandEight := 0
	HandSeven := 0
	HandSix := 0
	HandFive := 0
	HandFour := 0
	HandThree := 0
	HandTwo := 0

	DeckAce := 0
	DeckKing := 0
	DeckQueen := 0
	DeckJack := 0
	DeckT := 0
	DeckNine := 0
	DeckEight := 0
	DeckSeven := 0
	DeckSix := 0
	DeckFive := 0
	DeckFour := 0
	DeckThree := 0
	DeckTwo := 0

	// iterate over the hand and identify the ranks and their quantity
	for i := 0; i < len(hand); i++ {
		HandAce += strings.Count(hand[i].Rank, "A")
		HandKing += strings.Count(hand[i].Rank, "K")
		HandQueen += strings.Count(hand[i].Rank, "Q")
		HandJack += strings.Count(hand[i].Rank, "J")
		HandT += strings.Count(hand[i].Rank, "T")
		HandNine += strings.Count(hand[i].Rank, "9")
		HandEight += strings.Count(hand[i].Rank, "8")
		HandSeven += strings.Count(hand[i].Rank, "7")
		HandSix += strings.Count(hand[i].Rank, "6")
		HandFive += strings.Count(hand[i].Rank, "5")
		HandFour += strings.Count(hand[i].Rank, "4")
		HandThree += strings.Count(hand[i].Rank, "3")
		HandTwo += strings.Count(hand[i].Rank, "2")
	}

	rankHand := []int{HandAce, HandKing, HandQueen, HandJack, HandT, HandNine, HandEight, HandSeven, HandSix, HandFive, HandFour, HandThree, HandTwo}

	// highest frequency on top
	sort.Sort(sort.Reverse(sort.IntSlice(rankHand)))
	deckExchange = rankHand[0]

	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckAce += strings.Count(deck[i].Rank, "A")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckKing += strings.Count(deck[i].Rank, "K")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckQueen += strings.Count(deck[i].Rank, "Q")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckJack += strings.Count(deck[i].Rank, "J")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckT += strings.Count(deck[i].Rank, "T")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckNine += strings.Count(deck[i].Rank, "9")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckEight += strings.Count(deck[i].Rank, "8")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckSeven += strings.Count(deck[i].Rank, "7")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckSix += strings.Count(deck[i].Rank, "6")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckFive += strings.Count(deck[i].Rank, "5")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckFour += strings.Count(deck[i].Rank, "4")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckThree += strings.Count(deck[i].Rank, "3")
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckTwo += strings.Count(deck[i].Rank, "2")
	}

	// Total number of a certain Rank
	totA := HandAce + DeckAce
	totK := HandKing + DeckKing
	totQ := HandQueen + DeckQueen
	totJ := HandJack + DeckJack
	totT := HandT + DeckT
	totNine := HandNine + DeckNine
	totEight := HandEight + DeckEight
	totSeven := HandSeven + DeckSeven
	totSix := HandSix + DeckSix
	totFive := HandFive + DeckFive
	totFour := HandFour + DeckFour
	totThree := HandThree + DeckThree
	totTwo := HandTwo + DeckTwo

	rankTotal := []int{totA, totK, totQ, totJ, totT, totNine, totEight, totSeven, totSix, totFive, totFour, totThree, totTwo}

	// highest frequency on top
	sort.Sort(sort.Reverse(sort.IntSlice(rankTotal)))

	finalOutcome := -1
	firstOutcome := rankTotal[0]
	secondOutcome := rankTotal[1]

	// four-of-a-kind
	if firstOutcome == 4 {
		finalOutcome = 2
	}

	// full-house
	if firstOutcome == 3 && rankHand[0] > 1 {
		finalOutcome = 3
	}

	// three-of-a-kind
	if firstOutcome == 3 && rankHand[0] == 1 {
		finalOutcome = 6
	}

	// two-pairs
	if firstOutcome == 2 && secondOutcome == 2 {
		finalOutcome = 7
	}

	// one-pair
	if firstOutcome == 2 && rankHand[0] == 1 {
		finalOutcome = 8
	}
	if firstOutcome == 1 {
		finalOutcome = 9
	}
	return finalOutcome
}
