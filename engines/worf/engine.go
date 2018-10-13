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
		s := e.negamax(b, e.depth, c, m)
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove, bestScore
}

func (e *worf) negamax(b *Board, depth int, c Color, m Move) float64 {

	//if w := b.WithMove(m).Winner(); w != 0 {
	//	return Inf(w * c)
	//}

	b.AssertValid()

	if b.At(m.Dst) == WK {
		return Inf(-c * White)
	}
	if b.At(m.Dst) == BK {
		return Inf(-c * Black)
	}

	if depth == 0 {
		return e.Heuristic2(b, c, m)
	}

	value := Inf(1)

	b2 := b.WithMove(m)
	b = nil
	for _, m := range AllMoves(b2, -c) {
		v := e.negamax(b2, depth-1, -c, m) * -1
		value = min(value, v)
	}
	return value
}

type Heuristic func(*Board, Color, Move) float64

func (e *worf) Heuristic2(b *Board, c Color, m Move) float64 {
	NumEvals++

	if w := b.WithMove(m).Winner(); w != 0 {
		return Inf(w * c)
	}

	b = b.WithMove(m)
	h := 0.0
	for _, p := range b {
		h += valueOf[p+6]
	}
	return float64(c) * (h + e.noise())
}

var valueOf = [13]float64{
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

func (e *worf) noise() float64 {
	return e.rnd.Float64() / 1024
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
