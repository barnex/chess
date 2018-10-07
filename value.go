package chess

type Value struct {
	Win, Lose bool
	InMoves   int
	Heuristic float64
}

func (a Value) GT(b Value) bool {
	switch {
	case a.Win && !b.Win:
		return true
	case !a.Win && b.Win:
		return false
	case a.Lose && !b.Lose:
		return false
	case !a.Lose && b.Lose:
		return true
	case a.Win && b.Win:
		return a.InMoves < b.InMoves
	case a.Lose && b.Lose:
		return a.InMoves > b.InMoves
	default:
		return a.Heuristic > b.Heuristic
	}
}
