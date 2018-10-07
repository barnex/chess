package chess

import (
	"math/rand"
)

// EWorf thinks 1 move ahead.
func EWorf() Engine {
	return &eWorf{newRand()}
}

type eWorf struct {
	rnd *rand.Rand
}

func (e *eWorf) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = scoreMove(b, bestMove, c)
	)
	for _, m := range moves {
		if s := scoreMove(b, m, c); s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove
}

func scoreMove(b *Board, m Move, c Color) float64 {
	b2 := b.Copy()
	b2.Move(m.Src, m.Dst)
	return scoreBoard(b2, c)
}

func scoreBoard(b *Board, c Color) float64 {
	return float64(b.Winner() * c)
}
