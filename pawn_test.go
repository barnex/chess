package chess

func ExampleRulesBP() {
	b := Upright(&Board{
		bR, bN, bB, bQ, bK, bB, bN, bR,
		bP, bP, bP, 00, bP, bP, bP, bP,
		00, wP, 00, 00, 00, 00, bB, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, wP, 00, 00, 00, 00,
		00, 00, 00, bP, 00, 00, 00, 00,
		wP, 00, wP, wP, wP, wP, wP, wP,
		wR, wN, wB, wQ, wK, wB, wN, wR,
	})

	MarkAllMoves(b, bP)

	//Output:
	// moves for ♟d3
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · x ♙ x ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟a7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 x x · · · · ♝ ·
	// 5 x · · · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟b7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟c7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · x x · · · ♝ ·
	// 5 · · x · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟e7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · x · ♝ ·
	// 5 · · · · x · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟f7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · x ♝ ·
	// 5 · · · · · x · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟g7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♟h7
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ x
	// 5 · · · · · · · x
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h

}

func ExampleRulesWP() {
	b := Upright(&Board{
		bR, bN, bB, bQ, bK, bB, bN, bR,
		bP, bP, bP, 00, bP, bP, bP, bP,
		00, wP, 00, 00, 00, 00, bB, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, wP, 00, 00, 00, 00,
		00, 00, 00, bP, 00, 00, 00, 00,
		wP, 00, wP, wP, wP, wP, wP, wP,
		wR, wN, wB, wQ, wK, wB, wN, wR,
	})

	MarkAllMoves(b, wP)

	//Output:
	// moves for ♙a2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 x · · ♙ · · · ·
	// 3 x · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙c2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · x ♙ · · · ·
	// 3 · · x x · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙d2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙e2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ x · · ·
	// 3 · · · x x · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙f2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · x · ·
	// 3 · · · ♟ · x · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙g2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · x ·
	// 3 · · · ♟ · · x ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙h2
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · · x
	// 3 · · · ♟ · · · x
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙d4
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 ♟ ♟ ♟ · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · x · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
	//
	// moves for ♙b6
	// 8 ♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
	// 7 x ♟ x · ♟ ♟ ♟ ♟
	// 6 · ♙ · · · · ♝ ·
	// 5 · · · · · · · ·
	// 4 · · · ♙ · · · ·
	// 3 · · · ♟ · · · ·
	// 2 ♙ · ♙ ♙ ♙ ♙ ♙ ♙
	// 1 ♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	//   a b c d e f g h
}