package picard

import (
	"testing"

	. "github.com/barnex/chess"
)

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
