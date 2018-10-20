package picard

import (
	"fmt"

	. "github.com/barnex/chess"
)

func ExampleTactical_1() {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, BP, 00, BQ, 00, 00, 00,
		00, 00, 00, WP, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})

	e := New(0)
	PrintBestMove(e, b, White)

	//Output:
	// 8 · · · · · · · ·
	// 7 · · · · · · · ·
	// 6 · · · · · · · ·
	// 5 · · · · · · · ·
	// 4 · · ♟ · ♛ · · ·
	// 3 · · · ♙ · · · ·
	// 2 · · · · · · · ·
	// 1 · · · · · · · ·
	//   a b c d e f g h
	// d3e4

}

func PrintBestMove(e Engine, b *Board, c Color) {
	fmt.Println(b)

	m, _ := e.Move(b, c)
	fmt.Println(m)
}
