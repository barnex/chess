package picard

import (
	"testing"

	. "github.com/barnex/chess"
)

// Requires zero lookahead.
func TestTactical0(t *testing.T) {
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
	CheckBestMove(t, New(0), b, White, "e4")
}

// Requires 1 lookahead.
func TestTactical1(t *testing.T) {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, BP, 00, 00,
		00, 00, BP, 00, BN, 00, 00, 00,
		00, 00, 00, WB, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})
	CheckBestMove(t, New(1), b, White, "Bc4")
}

// Requires 1 lookahead, play as black
func TestTactical1B(t *testing.T) {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, WB, 00, 00,
		00, 00, WP, 00, WN, 00, 00, 00,
		00, 00, 00, BB, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})
	CheckBestMove(t, New(1), b, Black, "Bc4")
}

// White must capture the black king.
// Tests fails if AlphaBeta does not return immediately when a king is captured.
// Here, the white king would be captured back if the game had not ended.
func TestTacticalTakeTheKing(t *testing.T) {
	b := Upright(&Board{
		00, 00, BK, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, WK, 00, WQ, 00, BR, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})
	CheckBestMove(t, New(1), b, White, "Qc8")
}

func CheckBestMove(t *testing.T, e Engine, b *Board, c Color, want string) {
	t.Helper()
	m, _ := e.Move(b, c)
	w := MustParse(want, b, c)
	if m != w {
		t.Errorf("%v: have: %v, want: %v\n%v", c, m, w, b)
	}
}
