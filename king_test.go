package chess

func ExampleRulesK() {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, bP, 00, 00, 00, 00, 00,
		00, bK, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, bP, 00,
		00, bP, 00, 00, 00, wK, 00, 00,
		00, 00, 00, 00, 00, wP, wP, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})

	MarkAllMoves(b, wK, bK)

	//Output:
	// 	moves for ♔f3
	// 8 · · · · · · · ·
	// 7 · · ♟ · · · · ·
	// 6 · ♚ · · · · · ·
	// 5 · · · · · · · ·
	// 4 · · · · x x x ·
	// 3 · ♟ · · x ♔ x ·
	// 2 · · · · x ♙ ♙ ·
	// 1 · · · · · · · ·
	//   a b c d e f g h
	//
	// moves for ♚b6
	// 8 · · · · · · · ·
	// 7 x x ♟ · · · · ·
	// 6 x ♚ x · · · · ·
	// 5 x x x · · · · ·
	// 4 · · · · · · ♟ ·
	// 3 · ♟ · · · ♔ · ·
	// 2 · · · · · ♙ ♙ ·
	// 1 · · · · · · · ·
	//   a b c d e f g h

}