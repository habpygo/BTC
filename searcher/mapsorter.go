package searcher

import (
	"BTC/psychic/cards"
	"sort"
	"strings"
)

func MapSorter(pslice []string, length int) string {
	RankingMap := make(map[string]int)
	for i := 0; i < length; i++ {
		RankingMap[pslice[i]] = cards.SetStraightRank(pslice[i])
	}
	type pokerhand struct {
		Key   string
		Value int
	}
	var structslice []pokerhand
	for k, v := range RankingMap {
		structslice = append(structslice, pokerhand{k, v})
	}
	// sort
	sort.Slice(structslice, func(i, j int) bool {
		return structslice[i].Value > structslice[j].Value
	})

	var stringSlice []string
	for _, kv := range structslice {
		stringSlice = append(stringSlice, kv.Key)

	}
	stringSet := strings.Join(stringSlice, "")
	return stringSet
}

// take a straight map

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func SortMap(m map[string]int) PairList {
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}

	return pl
}
