package searcher

import (
	"BTC/psychic/cards"
	"sort"
	"strings"
)

const TotGame = 10

var MyCounter int

// LookForStraight searches for a straight
func LookForStraight(game []cards.Card) bool {

	type pokerhand struct {
		Key   string
		Value int
	}
	var ss []pokerhand
	var stringSlice []string
	type totGame []cards.Card
	var result = false

	RankedHand := make(map[string]int)

	// 1. SetRank voor game

	for i := 0; i < 10; i++ {
		MyCounter = i
		RankedHand[game[i].Rank] = cards.SetStraightRank(game[i].Rank) // powersetrankd
	}

	for key, value := range RankedHand {
		ss = append(ss, pokerhand{key, value})
	}

	// and sort it
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	// fmt.Println("the sorted ss is: ", ss) // [{a,v} {c,d} etc.]

	for _, kv := range ss {
		stringSlice = append(stringSlice, kv.Key)
	}
	//fmt.Println("stringSlice is: ", stringSlice)
	p := PowerSet(stringSlice)
	//fmt.Println("Powerset is: ", p)

	for i := len(p) - 1; i > 0; i-- {
		setp := strings.Join(p[i], "")
		if len(setp) == 5 {
			//	fmt.Println("setp is: ", setp)
			//	fmt.Println("StraightSet is: ", cards.StraightSet)
			if strings.Contains(cards.StraightSet, setp) {
				//		fmt.Println("The correct sequence is: ", setp)
				result = true
				break
			}
		}

		//	fmt.Println("result is: ", result)
	}
	return result
}

func copy_and_append_string(slice []string, elem string) []string {
	// wrong: return append(slice, elem)
	return append(append([]string(nil), slice...), elem)
}

func PowerSet(s []string) [][]string {
	if s == nil {
		return nil
	}
	r := [][]string{[]string{}}
	for _, es := range s {
		var u [][]string
		for _, er := range r {
			u = append(u, copy_and_append_string(er, es))
		}
		r = append(r, u...)
	}
	return r
}
