package chess

type Engine interface {
	Move(*Board, Color) (Move, float64)
}
