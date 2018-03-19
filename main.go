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

	file, noOfLinesInFile, err := utils.ReadLine("sampleinput")
	if err != nil {
		fmt.Println("couldn't read file")
	}
	fmt.Println("noOfLines is: ", noOfLinesInFile)

	for i := 0; i < noOfLinesInFile; i++ {

		for j := 0; j < len(file[i]); j = j + 3 {
			if j < len(file[i])-1 {
				card.Rank = string(file[i][j])
				card.Suit = string(file[i][j+1])
				Game = append(Game, card)
			}
		}
	}

	// play the game
	counter := 1
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

		FlushOutcome := searcher.LookForFlushLib(flushPotence1, numberfP1, Deck)
		if FlushOutcome == 0 {
			ResultString = "straight-flush"
			//GameArray = append(GameArray, ResultString)
			fmt.Println("YEAH ... straight-flush")
			//utils.FormatSolution(FlushOutcome)

		} else {
			ResultString = "flush"
			fmt.Println("just flush potential")
			ResultArray = append(ResultArray, ResultString)
		}
	}
}
