package riker

import (
	"testing"

	. "github.com/barnex/chess"
)

func BenchmarkMinimax2(b *testing.B) {
	e := New(2)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		e.Move(b, White)
	}
}

func BenchmarkMinimax3(b *testing.B) {
	e := New(3)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		e.Move(b, White)
	}
}
