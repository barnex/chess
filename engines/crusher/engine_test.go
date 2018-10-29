package crusher

import (
	"testing"

	"github.com/barnex/chess"
)

func TestAlphaBetaChange(t *testing.T) {

	const d = 4
	const N = 20
	var b1, b2 chess.Board
	{
		ea := New(d, 0)
		eb := New(d, 0)

		b1 = *game(ea, eb, N)
	}
	{
		ea := New(d, 0)
		ea.EnableAlphaBeta = true
		eb := New(d, 0)
		eb.EnableAlphaBeta = true

		b2 = *game(ea, eb, N)
	}
	if b1 != b2 {
		t.Errorf("have:\n%vwant:\n%v", b2.String(), b1.String())
	}

}

func game(ea, eb chess.Engine, moves int) *chess.Board {
	b := chess.NewBoard()
	for i := 0; i < moves; i++ {
		{
			m, _ := ea.Move(b, chess.White)
			b = b.WithMove(m)
		}
		{
			m, _ := eb.Move(b, chess.Black)
			b = b.WithMove(m)
		}
	}
	return b
}
