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

type kv struct {
	Key   string
	Value int
}

func SortMap(m map[string]int) []kv {

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})
	return ss
}

// 	p := make(PairList, len(m))
// 	i := 0
// 	for k, v := range m {
// 		p[i] = Pair{k, v}
// 		i++
// 	}
// 	sort.Sort(sort.Reverse(p))

// 	return p
// }

// func (p PairList) Len() int           { return len(p) }
// func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
