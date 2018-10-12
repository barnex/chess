package chess

import "math/rand"

type Heuristic func(*Board, Color) float64

func Heuristic0(_ *Board, _ Color) float64 {
	return noise()
}

func Heuristic1(b *Board, c Color) float64 {
	return float64(c) * (Inf(c*b.Winner()) + noise())
}

func Heuristic2(b *Board, c Color) float64 {
	h := 0.0
	for _, p := range b {
		h += valueOf[p+6]
	}
	return float64(c) * (Inf(b.Winner()) + h + noise())
}

var valueOf = [13]float64{
	wP + 6: 1,
	wN + 6: 5,
	wB + 6: 5,
	wR + 6: 10,
	wQ + 6: 15,
	bP + 6: -1,
	bN + 6: -5,
	bB + 6: -5,
	bR + 6: -10,
	bQ + 6: -15,
}

func noise() float64 {
	return rand.Float64() / 1024
}
