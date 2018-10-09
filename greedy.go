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

func (e *greedy) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = MinusInf
	)
	for _, m := range moves {
		b2 := b.WithMove(m)
		s := e.heuristicValue(b2)
		if s.GT(bestScore) {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove
}

func (e *greedy) heuristicValue(b *Board) Value {
	return Value{Win: b.Winner(), Heuristic: e.h(b)}
}
