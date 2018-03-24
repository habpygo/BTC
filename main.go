package main

import (
	"BTC/psychic/cards"
	"BTC/psychic/searcher"
	"BTC/psychic/utils"
	"fmt"
	"strings"
)

func main() {

	var ResultNumber int
	// file contains game in memory
	card := cards.Card{}
	Game := []cards.Card{}
	var Hand = Game[:]
	var Deck = Game[:]

	file, noOfLinesInFile, err := utils.ReadLine("sampleinput")
	if err != nil {
		fmt.Println("couldn't read file")
	}
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
		//fmt.Printf("%v\n", Game[h:h+10])
		fmt.Printf("%v%v %v%v %v%v %v%v %v%v %v%v %v%v %v%v %v%v %v%v\n", Game[h].Rank, Game[h].Suit, Game[h+1].Rank, Game[h+1].Suit, Game[h+2].Rank, Game[h+2].Suit, Game[h+3].Rank, Game[h+3].Suit, Game[h+4].Rank, Game[h+4].Suit, Game[h+5].Rank, Game[h+5].Suit, Game[h+6].Rank, Game[h+6].Suit, Game[h+7].Rank, Game[h+7].Suit, Game[h+8].Rank, Game[h+8].Suit, Game[h+9].Rank, Game[h+9].Suit)
	}
	fmt.Println("\n------------------------SAMPLE OUTPUT-----------------------------")

	// play the game
	counter := 0

	for k := 5; k+10 < 96; k = k + 10 { // start main loop
		Hand = Game[k-5 : k] // 0 - 5; 10 - 15; 20 - 25 etc
		Deck = Game[k : k+5] // 5 - 10; 15 - 25; 25 - 35 etc

		counter++

		// we look for highest possible rank, i.e. straight-flush and flush first
		flushPotence1, numberfP1, _, _ := searcher.LookForSuits(Hand)

		ResultNumber = searcher.FindRanks(Hand, Deck)

		StraightFlush := searcher.LookForFlushInDeck(flushPotence1, numberfP1, Deck)
		if StraightFlush {
			ResultNumber = 1
			searcher.Flush = false
			// OUTCOME
			fmt.Println("Hand: ", strings.Trim(fmt.Sprint(Hand[0]), "{}"), strings.Trim(fmt.Sprint(Hand[1]), "{}"), strings.Trim(fmt.Sprint(Hand[2]), "{}"), strings.Trim(fmt.Sprint(Hand[3]), "{}"), strings.Trim(fmt.Sprint(Hand[4]), "{}"), "Deck: ", strings.Trim(fmt.Sprint(Deck[0]), "{}"), strings.Trim(fmt.Sprint(Deck[1]), "{}"), strings.Trim(fmt.Sprint(Deck[2]), "{}"), strings.Trim(fmt.Sprint(Deck[3]), "{}"), strings.Trim(fmt.Sprint(Deck[4]), "{}"), cards.EvaluateResults(ResultNumber))
			//fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
			continue // we found the highest value and can skip the rest
		}

		if ResultNumber == 3 && !searcher.Flush { // we found full house, we can skip the rest
			fmt.Println("Hand: ", strings.Trim(fmt.Sprint(Hand[0]), "{}"), strings.Trim(fmt.Sprint(Hand[1]), "{}"), strings.Trim(fmt.Sprint(Hand[2]), "{}"), strings.Trim(fmt.Sprint(Hand[3]), "{}"), strings.Trim(fmt.Sprint(Hand[4]), "{}"), "Deck: ", strings.Trim(fmt.Sprint(Deck[0]), "{}"), strings.Trim(fmt.Sprint(Deck[1]), "{}"), strings.Trim(fmt.Sprint(Deck[2]), "{}"), strings.Trim(fmt.Sprint(Deck[3]), "{}"), strings.Trim(fmt.Sprint(Deck[4]), "{}"), cards.EvaluateResults(ResultNumber))
			//fmt.Printf("\nHand: %v Deck: %v Best Hand: %v", Hand, Deck, cards.EvaluateResults(ResultNumber))
			continue
		}

		if ResultNumber >= 3 {
			if searcher.Flush {
				ResultNumber = 4
			}
			searcher.Flush = false
			Straight := searcher.IsStraight(Hand, Deck)
			if Straight {
				ResultNumber = 5
			}
		}

		// OUTCOME
		//fmt.Printf("\nHand: %v %v Deck: %s Best Hand: %s", strings.Trim(fmt.Sprint(Hand[0]), "{}")),  Hand[1], Deck, cards.EvaluateResults(ResultNumber))
		fmt.Println("Hand: ", strings.Trim(fmt.Sprint(Hand[0]), "{}"), strings.Trim(fmt.Sprint(Hand[1]), "{}"), strings.Trim(fmt.Sprint(Hand[2]), "{}"), strings.Trim(fmt.Sprint(Hand[3]), "{}"), strings.Trim(fmt.Sprint(Hand[4]), "{}"), "Deck: ", strings.Trim(fmt.Sprint(Deck[0]), "{}"), strings.Trim(fmt.Sprint(Deck[1]), "{}"), strings.Trim(fmt.Sprint(Deck[2]), "{}"), strings.Trim(fmt.Sprint(Deck[3]), "{}"), strings.Trim(fmt.Sprint(Deck[4]), "{}"), cards.EvaluateResults(ResultNumber))

	}
}
