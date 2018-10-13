package chess

import (
	"math/rand"
)

func Greedy(h Heuristic) Engine {
	return &greedy{newRand(), h}
}

type greedy struct {
	rnd *rand.Rand
	h   Heuristic
}

func (e *greedy) Move(b *Board, c Color) (Move, float64) {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = Inf(-1)
	)
	for _, m := range moves {
		b2 := b.WithMove(m)
		s := e.h(b2, c)
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove, bestScore
}
