package chess

import (
	"fmt"
	"math/rand"
	"time"
)

func ERandom() Engine {
	return &eRandom{rand.New(rand.NewSource(time.Now().UnixNano()))}
}

type eRandom struct {
	rnd *rand.Rand
}

func (e *eRandom) Move(b *Board, c Color) Move {
	var pcs []Pos
	for p := (Pos{}); p.Valid(); p = p.Next() {
		if b.At(p).Color() == c {
			pcs = append(pcs, p)
		}
	}
	e.rnd.Shuffle(len(pcs), func(i, j int) { pcs[i], pcs[j] = pcs[j], pcs[i] })
	moves := make([]Pos, 64)
	for _, p := range pcs {
		moves = moves[:0]
		Moves(b, p, &moves)
		if len(moves) > 0 {
			return Move{p, moves[rand.Intn(len(moves))]}
		}
	}
	panic(fmt.Sprintf("no moves for %v:\n%v", c, b))
}
