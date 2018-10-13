package chess

import (
	"fmt"
	"math/rand"
	"time"
)

func Minimax(depth int, h Heuristic) Engine {
	return &minimax{newRand(), depth, h}
}

type minimax struct {
	rnd   *rand.Rand
	depth int
	h     Heuristic
}

var numEval int

func (e *minimax) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = Inf(-1)
	)
	numEval = 0
	start := time.Now()
	for _, m := range moves {
		b2 := b.WithMove(m)
		s := e.negamax(b2, e.depth, c)
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	d := time.Since(start)
	fmt.Println(numEval, "boards evaluated in", d, float64(numEval)/d.Seconds(), "/s")
	return bestMove
}

func (e *minimax) negamax(b *Board, depth int, c Color) float64 {
	if depth == 0 {
		numEval++
		return e.h(b, c)
	}

	counterMoves := allMoves(b, -c)

	value := Inf(1)
	for _, m := range counterMoves {
		b2 := b.WithMove(m)
		v := e.negamax(b2, depth-1, -c) * -1
		value = Min(value, v)
	}
	return value
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
