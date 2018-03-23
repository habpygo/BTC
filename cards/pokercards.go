package cards

type Card struct {
	Rank string
	Suit string
}

// to compare slices and subslices
var StringSet = "AKQJT98765432"
var StraightSet = "KQJT98765432A"

func SetRank(rank string) int {
	//RankMap := make(map[string]int)
	switch rank {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	default:
		return 0
		//panic("Unrecognized Card, someone is cheating?")

	}
	return 0
}

func SetStraightRank(rank string) int {
	//RankMap := make(map[string]int)
	switch rank {
	case "A":
		return 1
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	default:
		return 0
		//panic("Unrecognized Card, someone is cheating?")

	}
	return 0
}
