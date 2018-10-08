package chess

type Heuristic func(b *Board) float64

func Heuristic0(b *Board) float64 {
	return 0
}

func Heuristic1(b *Board) float64 {
	value := 0.0
	for _, p := range b {
		value += float64(p.Color())
	}
	return value
}
