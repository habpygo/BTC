package searcher

import (
	"BTC/psychic/cards"
	"fmt"
	"sort"
	"strings"
)

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

	// iterate over the hand and identify the ranks and their quantity
	for i := 0; i < len(hand); i++ {
		fmt.Println("hand[i].Rank for this round is: ", hand[i].Rank, "suit is: ", hand[i].Suit)
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

	DeckAce := FindTotalDeckRanks("A", HandAce, deck)
	DeckKing := FindTotalDeckRanks("K", HandKing, deck)
	DeckQueen := FindTotalDeckRanks("Q", HandQueen, deck)
	DeckJack := FindTotalDeckRanks("J", HandJack, deck)
	DeckT := FindTotalDeckRanks("T", HandT, deck)
	DeckNine := FindTotalDeckRanks("9", HandNine, deck)
	DeckEight := FindTotalDeckRanks("8", HandEight, deck)
	DeckSeven := FindTotalDeckRanks("7", HandSeven, deck)
	DeckSix := FindTotalDeckRanks("6", HandSix, deck)
	DeckFive := FindTotalDeckRanks("5", HandFive, deck)
	DeckFour := FindTotalDeckRanks("4", HandFour, deck)
	DeckThree := FindTotalDeckRanks("3", HandThree, deck)
	DeckTwo := FindTotalDeckRanks("2", HandTwo, deck)

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

	sort.Sort(sort.Reverse(sort.IntSlice(rankTotal)))
	firstOutcome := rankTotal[0]
	secondOutome := rankTotal[1]
	//thirdOutome := rankTotal[2]
	// fourthOutome := rankTotal[3]
	// fifthOutome := rankTotal[4]

	fmt.Println("The whole []int rankTotal is: ", rankTotal)
	fmt.Println("\nfirstOutcome is: ", firstOutcome)
	fmt.Println("secondOutcome is: ", secondOutome)

	// four-of-a-kind
	if firstOutcome == 4 {
		return 1
	}

	// full-house
	if firstOutcome == 3 && secondOutome >= 2 || firstOutcome >= 2 && secondOutome == 3 {
		return 2
	}

	// three-of-a-kind
	if firstOutcome == 3 {
		return 5
	}

	// two-pair
	if firstOutcome == 2 && secondOutome == 2 {
		return 6
	}

	// one-pair
	if firstOutcome == 2 && secondOutome == 0 || firstOutcome == 0 && secondOutome == 1 {
		return 7
	}
	if firstOutcome == 1 && secondOutome == 1 {
		return 8
	}
	return 8
}
