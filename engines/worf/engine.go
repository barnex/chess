package worf

/*

import (
	. "github.com/barnex/chess"
)

// New returns an engine with given depth.
func New(depth int) Engine {
	return &minimax{depth}
}

type minimax struct {
	depth int
}

type Node struct {
	board Board
	value float64
}

func (e *minimax) Move(b *Board, c Color) (Move, float64) {

	moves := AllMoves(b, c)

	var (
		bestMove  = Move{}
		bestScore = Inf(-1)
	)

	root := Node{
		board: *b,
		value: e.initialValue(b, c),
	}

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

func (e *minimax) negamax(n *Node, depth int, c Color) float64 {
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

func (e *minimax) initialValue(b *Board, c Color) float64 {

}
*/
