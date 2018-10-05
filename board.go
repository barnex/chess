package main

import (
	"bytes"
	"fmt"
)

type Board [64]Piece

func NewBoard() *Board {
	return &Board{
		WR, WN, WB, WQ, WK, WB, WN, WR,
		WP, WP, WP, WP, WP, WP, WP, WP,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		00, 00, 00, 00, 00, 00, 00, 00,
		BP, BP, BP, BP, BP, BP, BP, BP,
		BR, BN, BB, BQ, BK, BB, BN, BR,
	}
}

func (b *Board) At(p Pos) Piece {
	return b[p.Index()]
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for r := 7; r >= 0; r-- {
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
