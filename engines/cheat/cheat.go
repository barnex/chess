package cheat

import (
	. "github.com/barnex/chess"
)

// Cheat returns an engine that often makes illegal moves.
// Used for testing.
func New() Engine {
	return &cheat{}
}

type cheat struct{}

func (e *cheat) Move(b *Board, c Color) (Move, float64) {
	return Move{RC(0, 1), RC(3, 4)}, 0
}
