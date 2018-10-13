package worf

import (
	"math/rand"
	"time"

	. "github.com/barnex/chess"
)

func New(depth int) Engine {
	return &worf{newRand(), depth}
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
		s := float64(e.negamax(b, e.depth, c, m)) + e.noise()
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove, bestScore
}

func (e *worf) negamax(b *Board, depth int, c Color, m Move) int {

	if dst := b.At(m.Dst); dst == WK || dst == BK {
		return inf(-c * dst.Color())
	}

	if depth == 0 {
		return e.Heuristic2(b, c, m)
	}

	value := inf(1)

	b2 := b.WithMove(m)
	b = nil
	for _, m := range AllMoves(b2, -c) {
		v := e.negamax(b2, depth-1, -c, m) * -1
		value = min(value, v)
	}
	return value
}

func (e *worf) Heuristic2(b *Board, c Color, m Move) int {
	NumEvals++

	b = b.WithMove(m)
	h := 0
	for _, p := range b {
		h += valueOf[p+6]
	}
	return int(c) * h
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

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (e *worf) noise() float64 {
	return e.rnd.Float64() / 1024
}

func inf(c Color) int {
	return int(c) * 99999
}
