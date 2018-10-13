package riker

import (
	"math/rand"
	"time"

	. "github.com/barnex/chess"
)

// New returns an engine with given minimax depth
// and a value Heuristic2
func New(depth int) Engine {
	return &minimax{newRand(), depth, Heuristic2, 0}
}

type minimax struct {
	rnd     *rand.Rand
	depth   int
	h       Heuristic
	numEval int
}

func (e *minimax) Move(b *Board, c Color) (Move, float64) {
	moves := AllMoves(b, c)

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

	if w := b.Winner(); w != 0 {
		return Inf(w * c)
	}

	if depth == 0 {
		return e.h(b, c)
		e.numEval++
	}

	counterMoves := AllMoves(b, -c)

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

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
