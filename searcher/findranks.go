package searcher

import (
	"BTC/psychic/cards"
	"fmt"
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
		//fmt.Println("hand[i].Rank for this round is: ", hand[i].Rank, "suit is: ", hand[i].Suit)
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
	fmt.Println("rankHand is: ", rankHand)
	var topFive int
	for i := 0; i < NoOfCards; i++ {
		topFive += rankHand[i]
	}
	deckExchange = rankHand[0]

	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckAce += strings.Count(deck[i].Rank, "A")
		//fmt.Println("A to pick from Deck is: ", NoOfCards-HandAce, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckKing += strings.Count(deck[i].Rank, "K")
		//fmt.Println("K to pick from Deck is: ", NoOfCards-HandKing, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckQueen += strings.Count(deck[i].Rank, "Q")
		//fmt.Println("Q to pick from Deck is: ", NoOfCards-HandQueen, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckJack += strings.Count(deck[i].Rank, "J")
		//fmt.Println("J to pick from Deck is: ", NoOfCards-HandJack, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckT += strings.Count(deck[i].Rank, "T")
		//fmt.Println("T to pick from Deck is: ", NoOfCards-HandT, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckNine += strings.Count(deck[i].Rank, "9")
		//fmt.Println("9 to pick from Deck is: ", NoOfCards-HandNine, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckEight += strings.Count(deck[i].Rank, "8")
		//fmt.Println("8 to pick from Deck is: ", NoOfCards-HandEight, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckSeven += strings.Count(deck[i].Rank, "7")
		//fmt.Println("7 to pick from Deck is: ", NoOfCards-HandSeven, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckSix += strings.Count(deck[i].Rank, "6")
		//fmt.Println("6 to pick from Deck is: ", NoOfCards-HandSix, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckFive += strings.Count(deck[i].Rank, "5")
		//fmt.Println("5 to pick from Deck is: ", NoOfCards-HandFive, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckFour += strings.Count(deck[i].Rank, "4")
		//fmt.Println("4 to pick from Deck is: ", NoOfCards-HandFour, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckThree += strings.Count(deck[i].Rank, "3")
		//fmt.Println("3 to pick from Deck is: ", NoOfCards-HandThree, "Rank: ", deck[i].Rank)
	}
	for i := 0; i < NoOfCards-deckExchange; i++ {
		DeckTwo += strings.Count(deck[i].Rank, "2")
		//fmt.Println("2 to pick from Deck is: ", NoOfCards-HandTwo, "Rank: ", deck[i].Rank)
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
	// thirdOutome := rankTotal[2]
	// fourthOutome := rankTotal[3]
	// fifthOutcome := rankTotal[4]
	// cater for full-house

	fmt.Println("The whole []int rankTotal is: ", rankTotal)
	fmt.Println("\nfirstOutcome is: ", firstOutcome)
	fmt.Println("secondOutcome is: ", secondOutcome)

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

	// rankTotal[0] = 0
	// rankTotal[1] = 1
	return finalOutcome
}

// func checkThreeOfaK(firstOutcome, secondOutcome, thirdOutcome, fourthOutcome, fifthOutcome int) int {
// 	if firstOutcome == 1 && secondOutcome == 1 && thirdOutcome == 1 && fourthOutcome == 1 && fifthOutcome == 1 {
// 		secondOutcome--
// 	}
// }

//fmt.Println("rankHand is: ", rankHand)
//fmt.Println("=== deckExchange is: ", deckExchange)

// DeckAce := FindTotalDeckRanks("A", HandAce, deck)
// DeckKing := FindTotalDeckRanks("K", HandKing, deck)
// DeckQueen := FindTotalDeckRanks("Q", HandQueen, deck)
// DeckJack := FindTotalDeckRanks("J", HandJack, deck)
// DeckT := FindTotalDeckRanks("T", HandT, deck)
// DeckNine := FindTotalDeckRanks("9", HandNine, deck)
// DeckEight := FindTotalDeckRanks("8", HandEight, deck)
// DeckSeven := FindTotalDeckRanks("7", HandSeven, deck)
// DeckSix := FindTotalDeckRanks("6", HandSix, deck)
// DeckFive := FindTotalDeckRanks("5", HandFive, deck)
// DeckFour := FindTotalDeckRanks("4", HandFour, deck)
// DeckThree := FindTotalDeckRanks("3", HandThree, deck)
// DeckTwo := FindTotalDeckRanks("2", HandTwo, deck)
