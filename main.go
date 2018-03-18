package main

import (
	"BTC/psychic/searcher"
	"BTC/psychic/utils"
	"fmt"
)

func main() {

	GameArray, noOfLinesInFile, err := utils.ReadFile("sampleinput")
	for i := 0; i < noOfLinesInFile; i++ {
		var ResultString string

		if err != nil {
			fmt.Println("error in reading file")
		}
		fmt.Println("length of GameArray is: ", len(GameArray))
		// prepare Hand and Deck for each game i

		//fmt.Println(lineInput)
		Hand, Deck := utils.ReadHandAndDeck(GameArray, noOfLinesInFile)
		fmt.Println("GameArray is: ", GameArray)

		fmt.Printf("GAME no %x\n =================================\n", i)
		fmt.Println("Hand is: ", Hand)
		fmt.Println("Deck is: ", Deck)
		// we look for highest rank straight-flush and flush first
		flushPotence1, numberfP1, flushPotence2, numberfP2 := searcher.LookForSuits(Hand)
		fmt.Println("flushPotence1 is: ", flushPotence1, "and number of the same Suit are: ", numberfP1)
		fmt.Println("flushPotence1 is: ", flushPotence2, "and number of the same Suit are: ", numberfP2)

		flushOutcome := searcher.LookForFlush(flushPotence1, numberfP1, Deck)
		if flushOutcome == 0 {
			ResultString = "straight-flush"
			GameArray = append(GameArray, ResultString)
			fmt.Println("YEAH ... straight-flush")
		} else {
			ResultString = "flush"
			fmt.Println("just flush")
			GameArray = append(GameArray, ResultString)
		}
	}
}
