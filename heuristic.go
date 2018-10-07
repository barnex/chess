package chess

type Heuristic func(b *Board, c Color) float64

func Heuristic0(b *Board, c Color) float64 {
	return 0
}

func Heuristic1(b *Board, c Color) float64 {
	value := 0.0
	for _, p := range b {
		value += float64(p.Color() * c)
	}
	return value
}
