package worf

import (
	"math/rand"

	. "github.com/barnex/chess"
)

func New(depth int) Engine {
	return &worf{rand.New(rand.NewSource(123)), depth}
}

type worf struct {
	rnd   *rand.Rand
	depth int
}

func (e *worf) Move(b *Board, c Color) (Move, float64) {
	var (
		bestMove  = Move{}
		bestScore = Inf(-1)
	)
	for _, m := range AllMoves(b, c) {
		n := &Node{*b, Heuristic2(b, m)}
		s := float64(e.negamax(n, e.depth, c, m)) + e.noise() //break ties
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove, bestScore
}

type Node struct {
	board Board
	value int
}

func (n *Node) WithMove(m Move) *Node {
	return &Node{
		*n.board.WithMove(m),
		Heuristic2(&n.board, m),
	}
}

func (e *worf) negamax(n *Node, depth int, c Color, m Move) int {

	if dst := n.board.At(m.DstI()); dst == WK || dst == BK {
		return inf(-c * dst.Color())
	}

	if depth == 0 {
		return int(c) * Heuristic3(n, m)
	}

	value := inf(1)

	n2 := n.WithMove(m)
	n = nil
	for _, m := range AllMoves(&n2.board, -c) {
		v := e.negamax(n2, depth-1, -c, m) * -1
		value = min(value, v)
	}
	return value
}

func Heuristic3(n *Node, m Move) int {

	NumEvals++
	delta := -valueOf[n.board.At(m.DstI())+6]
	fast := (n.value + delta)

	return fast
}

func Heuristic2(b *Board, m Move) int {
	NumEvals++

	b = b.WithMove(m)
	h := 0
	for _, p := range b {
		h += valueOf[p+6]
	}
	return h
}

var valueOf = [13]int{
	WP + 6: 1,
	WN + 6: 6,
	WB + 6: 5,
	WR + 6: 10,
	WQ + 6: 20,
	BP + 6: -1,
	BN + 6: -6,
	BB + 6: -5,
	BR + 6: -10,
	BQ + 6: -20,
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (e *worf) noise() float64 {
	return e.rnd.Float64() / 1024
}

func inf(c Color) int {
	return int(c) * 99999
}
