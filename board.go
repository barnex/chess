package chess

import (
	"bytes"
	"fmt"
)

type Board [64]Piece

func NewBoard() *Board {
	return Upright(&Board{
		bR, bN, bB, bQ, bK, bB, bN, bR,
		bP, bP, bP, bP, bP, bP, bP, bP,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		wP, wP, wP, wP, wP, wP, wP, wP,
		wR, wN, wB, wQ, wK, wB, wN, wR,
	})
}

func Upright(b *Board) *Board {
	b2 := new(Board)
	for r := 0; r < 8; r++ {
		r2 := 7 - r
		for c := 0; c < 8; c++ {
			b2[RC(r2, c).Index()] = b[RC(r, c).Index()]
		}
	}
	return b2
}

func (b *Board) At(p Pos) Piece {
	return b[p.Index()]
}

func (b *Board) Move(src, dst Pos) {
	b[dst.Index()] = b[src.Index()]
	b[src.Index()] = 00
}

func (b *Board) Copy() *Board {
	c := new(Board)
	copy(c[:], b[:])
	return c
}

func (b *Board) Winner() Color {
	var haveWK, haveBK bool
	for _, p := range b {
		if p == wK {
			haveWK = true
		}
		if p == bK {
			haveBK = true
		}
	}
	switch {
	case !haveWK:
		return Black
	case !haveBK:
		return White
	default:
		return 0
	}
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for r := 7; r >= 0; r-- {
		fmt.Fprint(&buf, r+1)
		for c := 0; c < 8; c++ {
			fmt.Fprint(&buf, " ", b.At(RC(r, c)))
		}
		fmt.Fprintln(&buf)
	}
	fmt.Fprint(&buf, "  a b c d e f g h")
	return buf.String()
}
