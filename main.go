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

		//fmt.Println("\n=========== Game no: ", counter, "============")
		//fmt.Println("Hand: ", Hand)
		//fmt.Println("Deck: ", Deck)
		counter++

		// start analysis

		// we look for highest possible rank, i.e. straight-flush and flush first
		flushPotence1, numberfP1, flushPotence2, numberfP2 := searcher.LookForSuits(Hand)
		//fmt.Println("flushPotence1 is: ", flushPotence1, "and number of the same Suit are: ", numberfP1)
		//fmt.Println("flushPotence2 is: ", flushPotence2, "and number of the same Suit are: ", numberfP2)

		var StraightFlushOutcome bool
		var FlushOutcome bool
		var StraightOutcome bool
		resultNumber := -1

		StraightFlushOutcome, FlushOutcome = searcher.LookForFlushInDeck(flushPotence1, numberfP1, Deck)
		//fmt.Println("FlushOutcome is: ", StraightFlushOutcome, "so we have a straight-flush")

		StraightOutcome = searcher.LookForStraight(Hand, Deck)

		// loop over possibilities
		for resultNumber == -1 {

			if StraightFlushOutcome {
				resultNumber = 0
				fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(resultNumber))
				break // Done!
			}

			// return an int that depends on outcome
			resultNumber = searcher.FindRanks(Hand, Deck)

			value := resultNumber // fix the result outcome to avoid re-calculation
			//fmt.Println("From resultNumber in loop, value is: ", resultNumber)

			if resultNumber > 2 {
				_, FlushOutcome = searcher.LookForFlushInDeck(flushPotence2, numberfP2, Deck)
				if FlushOutcome {
					resultNumber = 3
					break
				} else {
					if StraightOutcome {
						resultNumber = 4
						//	fmt.Println("resultNumber is: ", resultNumber)
					}
				}
				resultNumber = value
			}
			fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(resultNumber))
			//	fmt.Println("resultNumber is: ", resultNumber)
		}
	} // End Game loop
}
