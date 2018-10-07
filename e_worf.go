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
		if s := scoreMove(b, m, c); s.GT(bestScore) {
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

func scoreMove(b *Board, m Move, c Color) Value {
	b2 := b.Copy()
	b2.Move(m.Src, m.Dst)
	return scoreBoard(b2, c)
}

func scoreBoard(b *Board, c Color) Value {
	w := b.Winner()
	if w == c {
		return Value{Win: true}
	}
	if w == -c {
		return Value{Lose: true}
	}
	return Value{Heuristic: Heuristic(b, c)}
}

func Heuristic(b *Board, c Color) float64 {
	return 0
}
