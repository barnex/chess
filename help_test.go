package chess

import (
	"fmt"
	"sort"
)

func MarkAllMoves(b *Board, typ ...Piece) {
	for _, typ := range typ {
		for i := range b {
			if b.At(Index(i)) == typ {
				fmt.Printf("moves for %v%v\n", b.At(Index(i)), Index(i).Pos())
				var moves []Move
				Moves(b, Index(i), &moves)
				mark(b, moves)
				fmt.Println()
			}
		}
	}
}

func mark(b *Board, p []Move) {
	c := b.copy()
	for _, p := range p {
		c[p.DstI()] = 100
	}
	fmt.Println(c)
}

func Sort(p []Pos) {
	sort.Slice(p, func(i, j int) bool { return p[i].String() < p[j].String() })
}
