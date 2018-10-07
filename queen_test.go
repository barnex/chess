package chess

func ExampleRulesQ() {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, bP, 00, 00, 00, 00, 00,
		00, bQ, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, bP, 00,
		00, bP, 00, 00, 00, wQ, 00, 00,
		00, 00, 00, 00, 00, wP, wP, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})

	MarkAllMoves(b, wQ, bQ)

	//Output:
	// moves for ♕f3
	// 8 x · · · · x · ·
	// 7 · x ♟ · · x · ·
	// 6 · ♛ x · · x · ·
	// 5 · · · x · x · ·
	// 4 · · · · x x x ·
	// 3 · x x x x ♕ x x
	// 2 · · · · x ♙ ♙ ·
	// 1 · · · x · · · ·
	//   a b c d e f g h
	//
	// moves for ♛b6
	// 8 · x · · · · · ·
	// 7 x x ♟ · · · · ·
	// 6 x ♛ x x x x x x
	// 5 x x x · · · · ·
	// 4 · x · x · · ♟ ·
	// 3 · ♟ · · x ♕ · ·
	// 2 · · · · · x ♙ ·
	// 1 · · · · · · · ·
	//   a b c d e f g h
}