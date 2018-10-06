package chess

import (
	"math/rand"
	"time"
)

func ERandom() Engine {
	return &eRandom{newRand()}
}

type eRandom struct {
	rnd *rand.Rand
}

func (e *eRandom) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)
	return moves[e.rnd.Intn(len(moves))]
}

func allMoves(b *Board, c Color) []Move {
	var moves []Move
	pos := make([]Pos, 0, 64)
	for p := (Pos{}); p.Valid(); p = p.Next() {
		if b.At(p).Color() == c {
			pos = pos[:0]
			Moves(b, p, &pos)
			for _, dst := range pos {
				moves = append(moves, Move{p, dst})
			}

		}
	}
	return moves
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
