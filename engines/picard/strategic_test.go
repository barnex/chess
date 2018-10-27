package picard

import (
	"testing"

	. "github.com/barnex/chess"
)

// Strategical weights set to mobility only.
// e4 allows most moves.
func TestMobility(t *testing.T) {
	b := Upright(&Board{
		BN, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, WP, WP, WP, WP, WP, 00,
		00, 00, WN, WB, WQ, WB, WN, 00,
	})
	e := NewOpts(2, false)
	e.Weight[0] = 0.001
	CheckBestMove(t, e, b, White, "e4")
}

// Strategic weights set to protection only.
func TestProtection(t *testing.T) {
	b := Upright(&Board{
		BP, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, WP, WP, WP, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})
	e := NewOpts(2, false)
	e.Weight[1] = 0.001
	CheckBestMove(t, e, b, White, "d3")
}

/*
// Strategic weights set to threat only.
func TestThreat(t *testing.T) {
	b := Upright(&Board{
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
	})
	e := NewOpts(2, false)
	e.Weight[1] = 0.001
	CheckBestMove(t, e, b, White, "d3")
}
*/
