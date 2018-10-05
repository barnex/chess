package main

import "fmt"

func ExampleNewBoard() {
	b := NewBoard()
	fmt.Println(b)

	//Output:
	//♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	//♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
	//· · · · · · · ·
	//· · · · · · · ·
	//· · · · · · · ·
	//· · · · · · · ·
	//♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
	//♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
}
