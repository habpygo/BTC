package cards

var ResultMap map[int]string

func EvaluateResults(result int) string {

	switch result {
	case 0:
		return "straight-flush"
	case 1:
		return "four-of-a-kind"
	case 2:
		return "full-house"
	case 3:
		return "flush"
	case 4:
		return "straight"
	case 5:
		return "three-of-a-kind"
	case 6:
		return "two-pairs"
	case 7:
		return "one-pair"
	case 8:
		return "highest-card"
	default:
		return "to sender"
	}
}
