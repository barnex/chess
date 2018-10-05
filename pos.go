package main

type Pos int

func (b *Board) At(p Pos) Piece {
	return b[p]
}
