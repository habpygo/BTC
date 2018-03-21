package cards

var ResultMap map[int]string

func EvaluateResults(result int) string {

	switch result {
	case 1:
		return "straight-flush"
	case 2:
		return "four-of-a-kind"
	case 3:
		return "full-house"
	case 4:
		return "flush"
	case 5:
		return "straight"
	case 6:
		return "three-of-a-kind"
	case 7:
		return "two-pairs"
	case 8:
		return "one-pair"
	case 9:
		return "highest-card"
	default:
		return "to sender"
	}
}
