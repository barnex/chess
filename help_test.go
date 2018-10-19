package chess

import (
	"fmt"
	"sort"
)

func MarkAllMoves(b *Board, typ ...Piece) {
	for _, typ := range typ {
		for p := RC(0, 0); p.Valid(); p = p.Next() {
			if b.At(p.Index()) == typ {
				fmt.Printf("moves for %v%v\n", b.At(p.Index()), p)
				var moves []Index
				Moves(b, p.Index(), &moves)
				mark(b, moves)
				fmt.Println()
			}
		}
	}
}

func mark(b *Board, p []Index) {
	c := b.copy()
	for _, p := range p {
		c[p] = 100
	}
	fmt.Println(c)
}

func Sort(p []Pos) {
	sort.Slice(p, func(i, j int) bool { return p[i].String() < p[j].String() })
}
