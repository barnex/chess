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
		00, 00, BP, 00, BK, 00, 00, 00,
		00, 00, 00, WQ, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})
	CheckBestMove(t, New(1), b, White, "Qc4")
}

func CheckBestMove(t *testing.T, e Engine, b *Board, c Color, want string) {
	t.Helper()
	m, _ := e.Move(b, c)
	w := MustParse(want, b, c)
	if m != w {
		t.Errorf("%v: have: %v, want: %v\n%v", c, m, w, b)
	}
}
