package searcher

import (
	"BTC/psychic/cards"
	"fmt"
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
		//fmt.Println("hand[i].Rank for this round is: ", hand[i].Rank)
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

	// iterate over the Deck depending on the number of Ranks found in Hand
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

	// map it to the counted value to sort on importance in the hand
	//handMap := map[string]int{"A": HandAce, "K": HandKing, "Q": HandQueen, "J": HandJack, "T": HandT, "9": HandNine, "8": HandEight, "7": HandSeven, "6": HandSix, "5": HandFive, "4": HandFour, "3": HandThree, "2": HandTwo}
	//deckMap := map[string]int{"A": DeckAce, "K": DeckKing, "Q": DeckQueen, "J": DeckJack, "T": DeckT, "9": DeckNine, "8": DeckEight, "7": DeckSeven, "6": DeckSix, "5": DeckFive, "4": DeckFour, "3": DeckThree, "2": DeckTwo}

	// totA := handMap["A"] + deckMap["A"]
	// totK := handMap["K"] + deckMap["K"]
	// totQ := handMap["Q"] + deckMap["Q"]
	// totJ := handMap["J"] + deckMap["J"]
	// totT := handMap["T"] + deckMap["T"]
	// totNine := handMap["9"] + deckMap["9"]
	// totEight := handMap["8"] + deckMap["8"]
	// totSeven := handMap["7"] + deckMap["7"]
	// totSix := handMap["6"] + deckMap["6"]
	// totFive := handMap["5"] + deckMap["5"]
	// totFour := handMap["4"] + deckMap["4"]
	// totThree := handMap["3"] + deckMap["3"]
	// totTwo := handMap["2"] + deckMap["2"]

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

	fmt.Println("HandA is: ", HandAce, "DeckAce is :", DeckAce, "and totA is: ", totA)
	fmt.Println("HandK is: ", HandKing, "DeckKing is :", DeckKing, "and totK is: ", totK)
	fmt.Println("HandQ is: ", HandQueen, "DeckQueen is :", DeckQueen, "and totQ is: ", totQ)
	fmt.Println("HandJ is: ", HandAce, "DeckJack is :", DeckJack, "and totJ is: ", totJ)
	fmt.Println("HandT is: ", HandT, "DeckT is :", DeckT, "and totT is: ", totT)
	fmt.Println("HandNine is: ", HandNine, "DeckNine is :", DeckNine, "and totNine is: ", totNine)
	fmt.Println("HandEight is: ", HandEight, "DeckEight is :", DeckEight, "and totEight is: ", totEight)
	fmt.Println("HandSeven is: ", HandSeven, "DeckSeven is :", DeckSeven, "and totSeven is: ", totSeven)
	fmt.Println("HandSix is: ", HandSix, "DeckSix is :", DeckSix, "and totSix is: ", totSix)
	fmt.Println("HandFive is: ", HandFive, "DeckFive is :", DeckFive, "and totFive is: ", totFive)
	fmt.Println("HandFour is: ", HandFour, "DeckFour is :", DeckFour, "and totFour is: ", totFour)
	fmt.Println("HandThree is: ", HandThree, "DeckThree is :", DeckThree, "and totThree is: ", totThree)
	fmt.Println("HandTwo is: ", HandTwo, "DeckTwo is :", DeckTwo, "and totTwo is: ", totTwo)

	// compare maps
	result := 0

	// 0 was straight-flush
	// 1 is four-of-a-kind, 2 is fullhouse, 3 is flush, 4 is straight, 5 is three-of-a-kind, 6 is two-pair, 7 is one-pair, 8 is highest-card
	return result
}

// n := map[int][]string{}
// 	var a []int
// 	var Rank []string
// 	var Freq []int
// sort the map
// for k, v := range m {
// 	n[v] = append(n[v], k)
// }

// for k := range n {
// 	a = append(a, k)
// }
// // we want the high frequencies on top and we want to pluck the first two
// // s is the Rank and k is the frequency
// sort.Sort(sort.Reverse(sort.IntSlice(a)))
// for _, k := range a {
// 	for _, s := range n[k] {
// 		fmt.Println("s is: ", s, "k is: ", k)

// 		//fmt.Printf("The outcome for Rank counting s is: %s, and k is %d\n", s, k)
// 	}
// }

// // return first highest ranks with their frequencies
