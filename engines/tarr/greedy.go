package tarr

import (
	"math/rand"
	"time"

	. "github.com/barnex/chess"
)

// New returns an engine that greedily takes the
// move with best heuristic, without thinking ahead.
func New(h Heuristic) Engine {
	return &greedy{newRand(), h}
}

type greedy struct {
	rnd *rand.Rand
	h   Heuristic
}

func (e *greedy) Move(b *Board, c Color) (Move, float64) {
	moves := AllMoves(b, c)

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

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
