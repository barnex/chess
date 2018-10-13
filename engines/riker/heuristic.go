package riker

import (
	"math/rand"

	. "github.com/barnex/chess"
)

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
