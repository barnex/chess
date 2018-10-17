package troi

import (
	"fmt"

	"github.com/barnex/chess"
)

func ExampleChange() {
	b := chess.NewBoard()

	e := New(2)

	c := chess.White
	for i := 0; i < 30; i++ {
		m, _ := e.Move(b, c)
		b = b.WithMove(m)
		c = -c
	}

	fmt.Println(b)

	//Output:
	// 8 · ♞ ♝ · ♚ ♝ ♞ ♜
	// 7 · ♟ · ♟ · ♟ ♟ ·
	// 6 · ♘ · · ♟ · · ·
	// 5 · ♗ ♟ · ♙ · · ♘
	// 4 · · · · · · · ♙
	// 3 · · ♛ · · · · ·
	// 2 ♖ ♙ · ♙ ♕ ♙ ♙ ·
	// 1 · · ♗ · · ♔ · ♖
	//   a b c d e f g h
}
