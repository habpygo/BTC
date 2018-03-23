package main

import (
	"BTC/psychic/cards"
	"BTC/psychic/searcher"
	"BTC/psychic/utils"
	"fmt"
)

func main() {

	var ResultNumber int
	//GameArray, noOfLinesInFile, err := utils.ReadFile("sampleinput")
	// file contains game in memory
	card := cards.Card{}
	Game := []cards.Card{}
	var Hand = Game[:]
	var Deck = Game[:]
	//var GameLine = Game[:]

	//
	file, noOfLinesInFile, err := utils.ReadLine("sampleinput")
	//file, noOfLinesInFile, err := utils.ReadLine("sampleinput-test")
	if err != nil {
		fmt.Println("couldn't read file")
	}
	//fmt.Println("noOfLinesFile is: ", noOfLinesInFile)

	// Read all the lines in: noOfLinesFile == 9, en(file[i]) == 29
	for i := 0; i < noOfLinesInFile; i++ {
		for j := 0; j < len(file[i]); j = j + 3 { // iterate +3 elements further
			if j < len(file[i])-1 {
				card.Rank = string(file[i][j])
				card.Suit = string(file[i][j+1])
				Game = append(Game, card) // perhaps we could build an interface around this
			}
		}
	}

	// Print out Games to play

	fmt.Println("\n------------------------SAMPLE INPUT------------------------\n")
	for h := 0; h < 90; h = h + 10 {
		fmt.Printf("%v\n", Game[h:h+10])
	}
	fmt.Println("\n--------------------------------------------SAMPLE OUTPUT--------------------------------------------")

	// play the game
	counter := 0

	for k := 5; k+10 < 96; k = k + 10 { // start main loop
		Hand = Game[k-5 : k] // 0 - 5; 10 - 15; 20 - 25 etc
		Deck = Game[k : k+5] // 5 - 10; 15 - 25; 25 - 35 etc

		// fmt.Println("\n=========== Game no: ", counter, "============")
		// fmt.Println("Hand: ", Hand)
		// fmt.Println("Deck: ", Deck)
		counter++

		// we look for highest possible rank, i.e. straight-flush and flush first
		flushPotence1, numberfP1, _, _ := searcher.LookForSuits(Hand)
		//fmt.Println("\nflushPotence1 is: ", flushPotence1, "and number of the same Suit are: ", numberfP1)
		//fmt.Println("\nflushPotence2 is: ", flushPotence2, "and number of the same Suit are: ", numberfP2)

		ResultNumber = searcher.FindRanks(Hand, Deck)
		//fmt.Println("ResultNumber in Main() is: ", ResultNumber)

		StraightFlush := searcher.LookForFlushInDeck(flushPotence1, numberfP1, Deck)
		if StraightFlush {
			ResultNumber = 1
			searcher.Flush = false
			// OUTCOME
			fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
			continue // we found the highest value and can skip the rest
		}

		if ResultNumber == 3 && !searcher.Flush { // we found full house, we can skip the rest
			fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
			continue
		}

		if ResultNumber >= 3 {
			if searcher.Flush {
				//fmt.Println(searcher.Flush)
				ResultNumber = 4
			}
			searcher.Flush = false
			Straight := searcher.IsStraight(Hand, Deck)
			if Straight {
				ResultNumber = 5
			}
		}

		// OUTCOME
		fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
	}
}
