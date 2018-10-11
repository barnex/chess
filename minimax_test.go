package chess

import "testing"

func BenchmarkMinimax2(b *testing.B) {
	e := Minimax(2, Heuristic2)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		e.Move(b, White)
	}
}

func BenchmarkMinimax3(b *testing.B) {
	e := Minimax(3, Heuristic2)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		e.Move(b, White)
	}
}
