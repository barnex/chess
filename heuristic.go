package chess

import "math/rand"

type Heuristic func(*Board, Color) Value

func Heuristic0(_ *Board, _ Color) Value {
	return Value{Heuristic: noise()}
}

func Heuristic1(b *Board, c Color) Value {
	return Value{Win: b.Winner(), Heuristic: noise()}.Mul(c)
}

func Heuristic2(b *Board, c Color) Value {
	h := 0.0
	for _, p := range b {
		h += valueOf[p]
	}
	return Value{Win: b.Winner(), Heuristic: h + noise()}.Mul(c)
}

var valueOf = map[Piece]float64{
	wP: 1,
	wN: 5,
	wB: 5,
	wR: 10,
	wQ: 15,

	bP: -1,
	bN: -5,
	bB: -5,
	bR: -10,
	bQ: -15,
}

func noise() float64 {
	return rand.Float64() / 1024
}
