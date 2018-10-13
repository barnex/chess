package chess

import (
	"math/rand"
)

func Minimax(depth int, h Heuristic) Engine {
	return &minimax{newRand(), depth, h, 0}
}

type minimax struct {
	rnd     *rand.Rand
	depth   int
	h       Heuristic
	numEval int
}

func (e *minimax) Move(b *Board, c Color) (Move, float64) {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = Inf(-1)
	)
	for _, m := range moves {
		b2 := b.WithMove(m)
		s := e.negamax(b2, e.depth, c)
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove, bestScore
}

func (e *minimax) negamax(b *Board, depth int, c Color) float64 {
	if depth == 0 {
		return e.h(b, c)
		e.numEval++
	}

	counterMoves := allMoves(b, -c)

	value := Inf(1)
	for _, m := range counterMoves {
		b2 := b.WithMove(m)
		v := e.negamax(b2, depth-1, -c) * -1
		value = min(value, v)
	}
	return value
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
