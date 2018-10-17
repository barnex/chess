package troi

import (
	"testing"

	. "github.com/barnex/chess"
)

func BenchmarkDepth2(b *testing.B) {
	e := New(2)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		e.Move(b, White)
	}
}

func BenchmarkDepth3(b *testing.B) {
	e := New(3)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		e.Move(b, White)
	}
}

func BenchmarkDepthGame3(b *testing.B) {
	e := New(3)
	for i := 0; i < b.N; i++ {
		b := NewBoard()
		for i := 0; i < 30; i++ {
			e.Move(b, White)
			e.Move(b, Black)
		}
	}
}
