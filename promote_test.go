package chess

import "fmt"

func ExamplePromotion() {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, WP, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, BP, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})

	b = b.WithMove(MoveP(RC(5, 1), RC(6, 1)))
	fmt.Println(b)
	b = b.WithMove(MoveP(RC(6, 1), RC(7, 1)))
	fmt.Println(b)

	b = b.WithMove(MoveP(RC(2, 4), RC(1, 4)))
	fmt.Println(b)
	b = b.WithMove(MoveP(RC(1, 4), RC(0, 4)))
	fmt.Println(b)

	//Output:
	// 8 · · · · · · · ·
	// 7 · ♙ · · · · · ·
	// 6 · · · · · · · ·
	// 5 · · · · · · · ·
	// 4 · · · · · · · ·
	// 3 · · · · ♟ · · ·
	// 2 · · · · · · · ·
	// 1 · · · · · · · ·
	//   a b c d e f g h
	// 8 · ♕ · · · · · ·
	// 7 · · · · · · · ·
	// 6 · · · · · · · ·
	// 5 · · · · · · · ·
	// 4 · · · · · · · ·
	// 3 · · · · ♟ · · ·
	// 2 · · · · · · · ·
	// 1 · · · · · · · ·
	//   a b c d e f g h
	// 8 · ♕ · · · · · ·
	// 7 · · · · · · · ·
	// 6 · · · · · · · ·
	// 5 · · · · · · · ·
	// 4 · · · · · · · ·
	// 3 · · · · · · · ·
	// 2 · · · · ♟ · · ·
	// 1 · · · · · · · ·
	//   a b c d e f g h
	// 8 · ♕ · · · · · ·
	// 7 · · · · · · · ·
	// 6 · · · · · · · ·
	// 5 · · · · · · · ·
	// 4 · · · · · · · ·
	// 3 · · · · · · · ·
	// 2 · · · · · · · ·
	// 1 · · · · ♛ · · ·
	//   a b c d e f g h

}
