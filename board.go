package chess

import (
	"bytes"
	"fmt"
)

type Board [64]Piece

func NewBoard() *Board {
	return Upright(&Board{
		BR, BN, BB, BQ, BK, BB, BN, BR,
		BP, BP, BP, BP, BP, BP, BP, BP,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		WP, WP, WP, WP, WP, WP, WP, WP,
		WR, WN, WB, WQ, WK, WB, WN, WR,
	})
}

func Upright(b *Board) *Board {
	b2 := new(Board)
	for r := 0; r < 8; r++ {
		r2 := 7 - r
		for c := 0; c < 8; c++ {
			b2[RC(r2, c).index()] = b.At(RC(r, c).Index())
		}
	}
	return b2
}

//func (b *Board) At(p Pos) Piece {
//	return b[int(p[0]<<3|p[1])]
//}

func (b *Board) At(p Index) Piece {
	return b[p]
}

func (b *Board) WithMove(m Move) *Board {

	// hack: always promote to queen
	// TODO: allow choice
	p := b.At(m.SrcI())
	if p == WP && m.Dst().Row() == 7 {
		p = WQ
	}
	if p == BP && m.Dst().Row() == 0 {
		p = BQ
	}

	c := b.copy()
	c[m.Dst().index()] = p
	c[m.Src().index()] = 00

	return c
}

func (b *Board) copy() *Board {
	c := new(Board)
	copy(c[:], b[:])
	return c
}

func (b *Board) Winner() Color {
	winner := Nobody
	for _, p := range b {
		if p == WK {
			winner += White
		}
		if p == BK {
			winner += Black
		}
	}
	return winner
}

func (b *Board) AssertValid() {
	var numWK, numBK int
	for _, p := range b {
		if p == WK {
			numWK++
		}
		if p == BK {
			numBK++
		}
	}
	if numWK > 1 || numBK > 1 || numWK+numBK < 1 {
		panic("invalid board\n" + b.String())
	}
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for r := 7; r >= 0; r-- {
		fmt.Fprint(&buf, r+1)
		for c := 0; c < 8; c++ {
			fmt.Fprint(&buf, " ", b.At(RC(r, c).Index()))
		}
		fmt.Fprintln(&buf)
	}
	fmt.Fprint(&buf, "  a b c d e f g h")
	return buf.String()
}
