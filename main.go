package main

import (
	"BTC/psychic/cards"
	"BTC/psychic/searcher"
	"BTC/psychic/utils"
	"fmt"
)

func main() {

	//GameArray, noOfLinesInFile, err := utils.ReadFile("sampleinput")
	// file contains game in memory
	card := cards.Card{}
	Game := []cards.Card{}
	var Hand = Game[:]
	var Deck = Game[:]
	var ResultArray []string
	var ResultString string

	//
	file, noOfLinesInFile, err := utils.ReadLine("sampleinput")
	if err != nil {
		fmt.Println("couldn't read file")
	}
	fmt.Println("noOfLinesFile is: ", noOfLinesInFile)

	// noOfLinesFile == 9
	// len(file[i]) == 29
	for i := 0; i < noOfLinesInFile; i++ {
		for j := 0; j < len(file[i]); j = j + 3 { // iterate +3 elements further
			if j < len(file[i])-1 {
				card.Rank = string(file[i][j])
				card.Suit = string(file[i][j+1])
				Game = append(Game, card) // perhaps we could build an interface around this
			}
		}
	}

	// play the game
	counter := 0
	for k := 5; k+10 < 96; k = k + 10 {
		Hand = Game[k-5 : k] // 0 - 5; 10 - 15; 20 - 25 etc
		Deck = Game[k : k+5] // 5 - 10; 15 - 25; 25 - 35 etc

		fmt.Println("\nGame no: ========== ", counter, "============")
		fmt.Println("Hand: ", Hand)
		fmt.Println("Deck: ", Deck)
		counter++
		// we look for highest rank straight-flush and flush first
		flushPotence1, numberfP1, flushPotence2, numberfP2 := searcher.LookForSuits(Hand)
		fmt.Println("flushPotence1 is: ", flushPotence1, "and number of the same Suit are: ", numberfP1)
		fmt.Println("flushPotence1 is: ", flushPotence2, "and number of the same Suit are: ", numberfP2)

		FlushOutcome := searcher.LookForFlushInDeck(flushPotence1, numberfP1, Deck)
		if FlushOutcome {
			ResultString = "straight-flush"
			fmt.Println("After analyzing we say: YEAH ... straight-flush. This is the highest we can get, so no further analysis warrented")
			ResultArray = append(ResultArray, ResultString) // TODO: make suitable array anc concatenate strings

		} else {
			firstRank, fRCount, secondRank, sRCount := searcher.FindRanks(Hand)
			fmt.Println("firstRank is :", firstRank, fRCount, "and secondRank is: ", sRCount, secondRank)
			// four-of-a-kind search here
			// full-house search here

			// then add flush to Result stack
			ResultString = "flush"
			fmt.Println("just flush potential, but we have to look for four-of-a-kind and or full-house first")
			ResultArray = append(ResultArray, ResultString)
		}

		// straight search here
		// three-of-kind search here
		// two-pairs search here
		// one-pair search here -- compare Hand with Deck with strings.Contains("deck[i].Rank, __the found in hand__ ")
		// highest-card search here

	}
}
