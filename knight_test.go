package chess

func ExampleRulesN() {
	b := Upright(&Board{
		00, 00, 00, bR, 00, 00, 00, 00,
		00, bN, 00, 00, 00, 00, 00, 00,
		00, 00, 00, wR, 00, 00, 00, 00,
		00, 00, wR, 00, wR, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, wN, 00, 00, 00, 00,
		00, 00, 00, 00, 00, bR, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})

	markAllMoves(b, wN, bN)

	//Output:
	// moves for ♘d3
	// 8 · · · ♜ · · · ·
	// 7 · ♞ · · · · · ·
	// 6 · · · ♖ · · · ·
	// 5 · · ♖ · ♖ · · ·
	// 4 · x · · · x · ·
	// 3 · · · ♘ · · · ·
	// 2 · x · · · x · ·
	// 1 · · x · x · · ·
	//   a b c d e f g h
	//
	// moves for ♞b7
	// 8 · · · ♜ · · · ·
	// 7 · ♞ · · · · · ·
	// 6 · · · x · · · ·
	// 5 x · x · ♖ · · ·
	// 4 · · · · · · · ·
	// 3 · · · ♘ · · · ·
	// 2 · · · · · ♜ · ·
	// 1 · · · · · · · ·
	//   a b c d e f g h

}
