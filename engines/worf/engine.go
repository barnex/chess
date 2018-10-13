package worf

import (
	"math/rand"
	"time"

	. "github.com/barnex/chess"
)

func New(depth int) Engine {
	return &minimax{newRand(), depth, Heuristic2, 0}
}

type minimax struct {
	rnd     *rand.Rand
	depth   int
	h       Heuristic
	numEval int
}

func (e *minimax) Move(b *Board, c Color) (Move, float64) {

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

func (e *minimax) negamax(b *Board, depth int, c Color, m Move) float64 {
	if depth == 0 {
		return e.h(b, c, m)
		e.numEval++
	}

	value := Inf(1)
	b = b.WithMove(m)
	for _, m := range AllMoves(b, -c) {
		v := e.negamax(b, depth-1, -c, m) * -1
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

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Heuristic func(*Board, Color, Move) float64

func Heuristic2(b *Board, c Color, m Move) float64 {
	NumEvals++

	dst := b.At(m.Dst)
	if dst == BK || dst == WK {
		return Inf(-dst.Color() * c)
	}

	b = b.WithMove(m)

	h := 0.0
	for _, p := range b {
		h += valueOf[p+6]
	}
	return float64(c) * (h + noise())
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

func noise() float64 {
	return rand.Float64() / 1024
}
