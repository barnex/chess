package main

import (
	"bytes"
	"fmt"
)

type Board [64]Piece

func NewBoard() *Board {
	return &Board{
		BR, BN, BB, BQ, BK, BB, BN, BR,
		BP, BP, BP, BP, BP, BP, BP, BP,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		WP, WP, WP, WP, WP, WP, WP, WP,
		WR, WN, WB, WQ, WK, WB, WN, WR,
	}
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			fmt.Fprint(&buf, b.At(RC(r, c)))
			if c < 7 {
				fmt.Fprint(&buf, " ")
			}
		}
		fmt.Fprintln(&buf)
	}
	return buf.String()
}

func (b *Board) At(p Pos) Piece {
	return b[p]
}

func (b *Board) Start() Pos { return 0 }
func (b *Board) End() Pos   { return Pos(len(b)) }
