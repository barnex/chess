package random

import (
	"math/rand"
	"time"

	. "github.com/barnex/chess"
)

// New returns an engine that makes random moves.
func New() Engine {
	return &random{newRand()}
}

type random struct {
	rnd *rand.Rand
}

func (e *random) Move(b *Board, c Color) (Move, float64) {
	moves := AllMoves(b, c)
	return moves[e.rnd.Intn(len(moves))], 0
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
