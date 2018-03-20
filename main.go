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

	//
	file, noOfLinesInFile, err := utils.ReadLine("sampleinput")
	if err != nil {
		fmt.Println("couldn't read file")
	}
	//fmt.Println("noOfLinesFile is: ", noOfLinesInFile)

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

		fmt.Println("\n=========== Game no: ", counter, "============")
		fmt.Println(Hand) // HAND
		fmt.Println(Deck)
		//fmt.Println("Hand: ", Hand)
		//fmt.Println("Deck: ", Deck)
		counter++

		// start analysis

		// we look for highest possible rank, i.e. straight-flush and flush first
		flushPotence1, numberfP1, flushPotence2, numberfP2 := searcher.LookForSuits(Hand)
		//fmt.Println("\nflushPotence1 is: ", flushPotence1, "and number of the same Suit are: ", numberfP1)
		fmt.Println("\nflushPotence2 is: ", flushPotence2, "and number of the same Suit are: ", numberfP2)

		var StraightFlush bool
		//var StraightOutcome bool
		ResultNumber := -1
		//ResultNumber2 := 0
		//var ResultArray []int

		StraightFlush = searcher.LookForFlushInDeck(flushPotence1, numberfP1, Deck)
		//fmt.Println("FlushOutcome is: ", StraightFlushOutcome, "so we have a straight-flush")

		//
		//StraightOutcome = searcher.LookForStraight(Hand, Deck)

		// loop over possibilities
		for ResultNumber == -1 {
			if StraightFlush {
				ResultNumber = 0
				//fmt.Println("From resultNumber in loop, value is: ", ResultNumber)
				//fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
				break // Done!
			}

			Straight := searcher.LookForStraight(Hand, Deck)
			if Straight {
				ResultNumber = 4
				//break
			}

			// now that we found the highest possible rank return an int that depends on outcome
			ResultNumber = searcher.FindRanks(Hand, Deck)
			//Value := ResultNumber
			//fmt.Println("\nFrom resultNumber in loop, value is: ", ResultNumber)

			// but don't forget the flush and straight
			// if ResultNumber > 2 {
			// 	ResultArray = append(ResultArray, ResultNumber)

			// 	StraightFlush = searcher.LookForFlushInDeck(flushPotence2, numberfP2, Deck)
			// 	if !StraightFlush { // it's a normal flush
			// 		ResultNumber2 = 3
			// 		ResultArray = append(ResultArray, ResultNumber2)
			// 		if ResultNumber < ResultNumber2 {
			// 		}

			// 	}
			// 	if StraightOutcome {
			// 		ResultNumber2 = 4
			// 		ResultArray = append(ResultArray, ResultNumber2)
			// 	}
			// 	sort.Ints(ResultArray)
			// 	fmt.Println("\nResultArray is: ", ResultArray)
			// 	ResultNumber = ResultArray[0]
			// }
		}

		fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
		//fmt.Println("\nresultNumber is: ", ResultNumber)
	}
} // End Game loop
