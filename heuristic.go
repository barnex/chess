package chess

type Heuristic func(b *Board) float64

func Heuristic0(b *Board) float64 {
	return 0
}

func Heuristic1(b *Board) float64 {
	value := 0.0
	for _, p := range b {
		value += valueOf[p]
	}
	return value
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
