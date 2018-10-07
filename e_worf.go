package chess

import (
	"math/rand"
)

// EWorf thinks 1 move ahead.
func EWorf(h Heuristic) Engine {
	return &eWorf{newRand(), h}
}

type Heuristic func(b *Board, c Color) float64

type eWorf struct {
	rnd *rand.Rand
	h   Heuristic
}

func (e *eWorf) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = e.scoreMove(b, bestMove, c)
	)
	for _, m := range moves {
		if s := e.scoreMove(b, m, c); s.GT(bestScore) {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove
}

//func scoreRec(d int, b *Board, m Move, c Color) float64 {
//	if d == 0 {
//		return scoreMove(b, m, c)
//	}
//
//}

func (e *eWorf) scoreMove(b *Board, m Move, c Color) Value {
	b2 := b.Copy()
	b2.Move(m.Src, m.Dst)
	return e.scoreBoard(b2, c)
}

func (e *eWorf) scoreBoard(b *Board, c Color) Value {
	w := b.Winner()
	if w == c {
		return Value{Win: true}
	}
	if w == -c {
		return Value{Lose: true}
	}
	return Value{Heuristic: e.h(b, c)}
}

func Heuristic0(b *Board, c Color) float64 {
	return 0
}

func Heuristic1(b *Board, c Color) float64 {
	value := 0.0
	for _, p := range b {
		value += float64(p.Color() * c)
	}
	return value
}
