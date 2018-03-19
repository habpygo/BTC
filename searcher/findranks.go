package searcher

import (
	"BTC/psychic/cards"
	"fmt"
	"sort"
	"strings"
)

func FindRanks(hand []cards.Card) (string, int, string, int) {

	Ace := 0
	King := 0
	Queen := 0
	Jack := 0
	T := 0
	Nine := 0
	Eight := 0
	Seven := 0
	Six := 0
	Five := 0
	Four := 0
	Three := 0
	Two := 0

	// iterate over the hand and identify the ranks and their quantity
	for i := 0; i < len(hand); i++ {
		fmt.Println("hand[i].Rank for this round is: ", hand[i].Rank)
		Ace += strings.Count(hand[i].Rank, "A")
		King += strings.Count(hand[i].Rank, "K")
		Queen += strings.Count(hand[i].Rank, "Q")
		Jack += strings.Count(hand[i].Rank, "J")
		T += strings.Count(hand[i].Rank, "T")
		Nine += strings.Count(hand[i].Rank, "9")
		Eight += strings.Count(hand[i].Rank, "8")
		Seven += strings.Count(hand[i].Rank, "7")
		Six += strings.Count(hand[i].Rank, "6")
		Five += strings.Count(hand[i].Rank, "5")
		Four += strings.Count(hand[i].Rank, "4")
		Three += strings.Count(hand[i].Rank, "3")
		Two += strings.Count(hand[i].Rank, "2")
	}

	// map it to the counted value to sort on importance in the hand
	m := map[string]int{"A": Ace, "K": King, "Q": Queen, "J": Jack, "T": T, "9": Nine, "8": Eight, "7": Seven, "6": Six, "5": Five, "4": Four, "3": Three, "2": Two}
	n := map[int][]string{}
	var a []int

	// sort the map
	for k, v := range m {
		n[v] = append(n[v], k)
	}

	for k := range n {
		a = append(a, k)
	}
	// we want the high frequencies on top and we want to pluck the first two
	// s is the Rank and k is the Suit
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		fmt.Println("k is: ", k)
		for _, s := range n[k] {
			fmt.Printf("The outcome for Rank counting s is: %s, and k is %d\n", s, k)
			fmt.Println("s[0] is: ", string(s[0]), "and k[0] is: ", k)
		}

	}
	// return first highest ranks with their frequencies
	return "", 0, "", 0
}
